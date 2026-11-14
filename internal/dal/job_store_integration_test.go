package dal

import (
	"database/sql"
	"errors"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"backend-lab-go/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// These tests use an isolated schema in the PostgreSQL database named by
// LEASE_TEST_DATABASE_URL. They never use the application's default DSN.
func leaseTestStore(t *testing.T) *JobStore {
	t.Helper()
	dsn := os.Getenv("LEASE_TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("set LEASE_TEST_DATABASE_URL to run PostgreSQL lease integration tests")
	}
	parsedConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		t.Fatalf("parse LEASE_TEST_DATABASE_URL: %v", err)
	}
	if parsedConfig.Database != "backend_lab_test" {
		t.Fatalf("refusing database %q: lease integration tests require database exactly named backend_lab_test", parsedConfig.Database)
	}

	admin, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}
	if err := admin.Ping(); err != nil {
		admin.Close()
		t.Fatal(err)
	}

	schema := "lease_test_" + strings.ReplaceAll(uuid.NewString(), "-", "")
	if _, err := admin.Exec(`CREATE SCHEMA ` + schema); err != nil {
		admin.Close()
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, dropErr := admin.Exec(`DROP SCHEMA ` + schema + ` CASCADE`)
		admin.Close()
		if dropErr != nil {
			t.Errorf("drop test schema: %v", dropErr)
		}
	})

	parsedConfig.RuntimeParams["search_path"] = schema
	db := stdlib.OpenDB(*parsedConfig)
	t.Cleanup(func() { db.Close() })
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	var currentSchema string
	if err := db.QueryRow(`SELECT current_schema()`).Scan(&currentSchema); err != nil {
		t.Fatal(err)
	}
	if currentSchema != schema {
		t.Fatalf("connection search_path resolved to %q, want isolated schema %q", currentSchema, schema)
	}
	_, err = db.Exec(`CREATE TABLE jobs (
		id BIGSERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending',
		ctx JSONB NOT NULL DEFAULT '{}'::jsonb,
		lease_token TEXT,
		lease_expires_at TIMESTAMPTZ,
		result JSONB,
		error_message TEXT
	)`)
	if err != nil {
		t.Fatal(err)
	}
	return NewJobStore(db)
}

func TestLeaseClaimRenewTakeoverAndFinish(t *testing.T) {
	store := leaseTestStore(t)
	created, err := store.Create("lease scenario", nil)
	if err != nil {
		t.Fatal(err)
	}

	workerA, err := store.ClaimNext(time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if workerA.ID != created.ID || workerA.LeaseToken == "" {
		t.Fatalf("unexpected claim: %+v", workerA)
	}
	if err := store.ExtendLease(workerA.ID, "wrong-token"); err == nil {
		t.Fatal("renewal with wrong token succeeded")
	}
	if err := store.ExtendLease(workerA.ID, workerA.LeaseToken); err != nil {
		t.Fatalf("valid renewal failed: %v", err)
	}

	if _, err := store.db.Exec(`UPDATE jobs SET lease_expires_at = NOW() - INTERVAL '1 second' WHERE id = $1`, workerA.ID); err != nil {
		t.Fatal(err)
	}
	workerB, err := store.ClaimNext(time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if workerB.LeaseToken == workerA.LeaseToken {
		t.Fatal("takeover did not rotate lease token")
	}
	if err := store.ExtendLease(workerA.ID, workerA.LeaseToken); err == nil {
		t.Fatal("expired worker renewed after takeover")
	}
	if err := store.UpdateStatus(workerA.ID, model.JobRunning, model.JobSucceeded, workerA.LeaseToken); err == nil {
		t.Fatal("expired worker finished after takeover")
	}
	current, err := store.GetJobByID(workerB.ID)
	if err != nil {
		t.Fatal(err)
	}
	if current.Status != model.JobRunning || current.LeaseToken != workerB.LeaseToken {
		t.Fatalf("old worker changed row after takeover: status=%q token=%q, want running with current token", current.Status, current.LeaseToken)
	}
	if err := store.UpdateStatus(workerB.ID, model.JobRunning, model.JobSucceeded, workerB.LeaseToken); err != nil {
		t.Fatalf("current lease owner could not finish: %v", err)
	}
	current, err = store.GetJobByID(workerB.ID)
	if err != nil {
		t.Fatal(err)
	}
	if current.Status != model.JobSucceeded {
		t.Fatalf("current worker finish did not persist: status=%q", current.Status)
	}
}

func TestExpiredLeaseCannotBeRenewedOrFinished(t *testing.T) {
	store := leaseTestStore(t)
	job, err := store.Create("expired lease", nil)
	if err != nil {
		t.Fatal(err)
	}
	claimed, err := store.ClaimNext(time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.db.Exec(`UPDATE jobs SET lease_expires_at = NOW() - INTERVAL '1 second' WHERE id = $1`, job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if err := store.ExtendLease(job.ID, claimed.LeaseToken); err == nil {
		t.Fatal("expired lease renewal succeeded")
	}
	if err := store.UpdateStatus(job.ID, model.JobRunning, model.JobSucceeded, claimed.LeaseToken); err == nil {
		t.Fatal("expired lease completion succeeded")
	}
}

func TestConcurrentClaimNextDoesNotDuplicateLease(t *testing.T) {
	store := leaseTestStore(t)
	created, err := store.Create("concurrent claim", nil)
	if err != nil {
		t.Fatal(err)
	}

	const claimantCount = 2
	start := make(chan struct{})
	var wg sync.WaitGroup
	jobs := make(chan model.Job, claimantCount)
	errs := make(chan error, claimantCount)
	for i := 0; i < claimantCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			job, err := store.ClaimNext(time.Minute)
			if err != nil {
				errs <- err
				return
			}
			jobs <- job
		}()
	}
	close(start)
	wg.Wait()
	close(jobs)
	close(errs)

	var claimed []model.Job
	for job := range jobs {
		claimed = append(claimed, job)
	}
	var noRows int
	for err := range errs {
		if errors.Is(err, sql.ErrNoRows) {
			noRows++
		} else {
			t.Fatalf("ClaimNext failed unexpectedly: %v", err)
		}
	}
	if len(claimed) != 1 || noRows != claimantCount-1 {
		t.Fatalf("claim results: successes=%d no_rows=%d, want 1 success and %d no_rows", len(claimed), noRows, claimantCount-1)
	}
	if claimed[0].ID != created.ID || claimed[0].LeaseToken == "" {
		t.Fatalf("unexpected claimed job: %+v", claimed[0])
	}

	stored, err := store.GetJobByID(created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if stored.Status != model.JobRunning || stored.LeaseToken != claimed[0].LeaseToken {
		t.Fatalf("database row does not match sole claim: status=%q token=%q", stored.Status, stored.LeaseToken)
	}
}
