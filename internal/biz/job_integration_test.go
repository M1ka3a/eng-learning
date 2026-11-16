package biz

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"backend-lab-go/internal/dal"
	"backend-lab-go/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// bizTestJobBiz builds a JobBiz backed by an isolated schema in the PostgreSQL
// database named by LEASE_TEST_DATABASE_URL. It applies the same safety guards
// as the DAL lease integration tests and never uses the application's DSN.
func bizTestJobBiz(t *testing.T) *JobBiz {
	t.Helper()
	dsn := os.Getenv("LEASE_TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("set LEASE_TEST_DATABASE_URL to run PostgreSQL biz integration tests")
	}
	parsedConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		t.Fatalf("parse LEASE_TEST_DATABASE_URL: %v", err)
	}
	if parsedConfig.Database != "backend_lab_test" {
		t.Fatalf("refusing database %q: integration tests require database exactly named backend_lab_test", parsedConfig.Database)
	}

	admin, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}
	if err := admin.Ping(); err != nil {
		admin.Close()
		t.Fatal(err)
	}

	schema := "biz_test_" + strings.ReplaceAll(uuid.NewString(), "-", "")
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
	var currentSchema string
	if err := db.QueryRow(`SELECT current_schema()`).Scan(&currentSchema); err != nil {
		t.Fatal(err)
	}
	if currentSchema != schema {
		t.Fatalf("connection search_path resolved to %q, want isolated schema %q", currentSchema, schema)
	}
	if _, err := db.Exec(`CREATE TABLE jobs (
		id BIGSERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending',
		ctx JSONB NOT NULL DEFAULT '{}'::jsonb,
		lease_token TEXT,
		lease_expires_at TIMESTAMPTZ,
		result JSONB,
		error_message TEXT
	)`); err != nil {
		t.Fatal(err)
	}
	return NewJobBiz(dal.NewJobStore(db))
}

// TestCreateClaimFinishSucceeds walks one ticket through the same Biz calls the
// worker makes: Create -> ClaimNext -> RouteTicket -> Finish.
func TestCreateClaimFinishSucceeds(t *testing.T) {
	jobBiz := bizTestJobBiz(t)

	ticket, err := json.Marshal(TicketDetails{
		UID:         "user-1",
		OrderID:     "order-1",
		Description: "charged twice",
		Category:    model.TicketPaymentIssue,
	})
	if err != nil {
		t.Fatal(err)
	}

	created, err := jobBiz.Create("ticket", ticket)
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if created.Status != model.JobPending {
		t.Fatalf("created status = %q, want %q", created.Status, model.JobPending)
	}

	claimed, err := jobBiz.ClaimNext()
	if err != nil {
		t.Fatalf("ClaimNext: %v", err)
	}
	if claimed.ID != created.ID {
		t.Fatalf("claimed job %d, want %d", claimed.ID, created.ID)
	}

	result, runErr := jobBiz.RouteTicket(claimed.Ctx)
	if runErr != nil {
		t.Fatalf("RouteTicket: %v", runErr)
	}
	if err := jobBiz.Finish(claimed.ID, claimed.LeaseToken, result, nil); err != nil {
		t.Fatalf("Finish: %v", err)
	}

	final, err := jobBiz.GetJobByID(context.Background(), created.ID)
	if err != nil {
		t.Fatalf("GetJobByID: %v", err)
	}
	if final.Status != model.JobSucceeded {
		t.Fatalf("final status = %q, want %q (lease_expires_at=%s)", final.Status, model.JobSucceeded, final.LeaseExpiresAt)
	}
}
