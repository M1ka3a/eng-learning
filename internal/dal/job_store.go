package dal

import (
	"backend-lab-go/internal/model"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type JobStore struct {
	db *sql.DB
}

func NewJobStore(db *sql.DB) *JobStore {
	return &JobStore{db: db}
}

func (s *JobStore) Create(title string, ctx json.RawMessage) (model.Job, error) {
	return s.CreateWithStatus(title, model.JobPending, ctx)
}

func (s *JobStore) CreateWithStatus(title string, status model.JobStatus, ctx json.RawMessage) (model.Job, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return model.Job{}, model.ErrJobTitleInvalid
	}

	ctxJSON := "{}"

	if len(ctx) > 0 {
		if !json.Valid(ctx) {
			return model.Job{}, model.ErrJobCtxInvalid
		}
		ctxJSON = string(ctx)
	}

	job := model.Job{}

	err := s.db.QueryRow(
		`INSERT INTO jobs (title, ctx, status)
		 VALUES ($1, $2, $3)
		 RETURNING id, title,status, ctx`,
		title, ctxJSON, status,
	).Scan(&job.ID, &job.Title, &job.Status, &job.Ctx)
	if err != nil {
		return model.Job{}, err
	}

	return job, nil
}

func (s *JobStore) GetJobByID(ctx context.Context, id int64) (model.Job, error) {
	return scanJob(s.db.QueryRowContext(ctx,
		`SELECT id, title, status, ctx, lease_token, lease_expires_at, result, error_message
         FROM jobs
         WHERE id = $1`,
		id,
	))
}

type rowScanner interface{ Scan(dest ...any) error }

func scanJob(row rowScanner) (model.Job, error) {
	var job model.Job
	var leaseToken sql.NullString
	var leaseExpiresAt sql.NullTime
	var result sql.NullString
	var errorMessage sql.NullString
	err := row.Scan(
		&job.ID,
		&job.Title,
		&job.Status,
		&job.Ctx,
		&leaseToken,
		&leaseExpiresAt,
		&result,
		&errorMessage,
	)

	if err != nil {
		return model.Job{}, err
	}
	if leaseToken.Valid {
		job.LeaseToken = leaseToken.String
	}
	if leaseExpiresAt.Valid {
		job.LeaseExpiresAt = leaseExpiresAt.Time.Format(time.RFC3339Nano)
	}
	if result.Valid {
		job.Result = json.RawMessage(result.String)
	}
	if errorMessage.Valid {
		job.ErrorMessage = errorMessage.String
	}

	return job, nil
}

func (s *JobStore) ListByUID(uid string) ([]model.Job, error) {
	rows, err := s.db.Query(
		`SELECT id, title, status, ctx, lease_token, lease_expires_at, result, error_message
		 FROM jobs WHERE ctx->>'uid' = $1 ORDER BY id DESC LIMIT 100`, uid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	jobs := make([]model.Job, 0)
	for rows.Next() {
		job, err := scanJob(rows)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, rows.Err()
}

func (s *JobStore) AddOrderID(ctx context.Context, id int64, uid, orderID string) (model.Job, error) {
	var job model.Job
	err := s.db.QueryRowContext(ctx,
		`UPDATE jobs
		 SET ctx = jsonb_set(ctx, '{order_id}', to_jsonb($3::text), true),
		     status = 'pending'
		 WHERE id = $1 AND status = 'needs_input' AND ctx->>'uid' = $2
		 RETURNING id, title, status, ctx`,
		id, uid, orderID,
	).Scan(&job.ID, &job.Title, &job.Status, &job.Ctx)
	if err == sql.ErrNoRows {
		return model.Job{}, model.ErrJobNotEditable
	}
	if err != nil {
		return model.Job{}, err
	}
	return job, nil
}

func (s *JobStore) ClaimNext(duration time.Duration) (model.Job, error) {
	job := model.Job{}
	leaseToken, err := newLeaseToken()
	if err != nil {
		return model.Job{}, err
	}

	err = s.db.QueryRow(
		`
		WITH candidate AS (
			SELECT id
			FROM jobs
			WHERE status = 'pending'
				OR (status = 'running' AND lease_expires_at <= NOW())
			ORDER BY id
			FOR UPDATE SKIP LOCKED
			LIMIT 1
		)
		UPDATE jobs AS j
		SET status = 'running',
			lease_token = $1,
			lease_expires_at = NOW() + ($2 * INTERVAL '1 second')
		FROM candidate
		WHERE j.id = candidate.id
		RETURNING j.id, j.title, j.status, j.ctx, j.lease_token
		`,
		leaseToken, duration.Seconds(),
	).Scan(
		&job.ID,
		&job.Title,
		&job.Status,
		&job.Ctx,
		&job.LeaseToken,
	)
	if err != nil {
		return model.Job{}, err
	}
	return job, nil
}

func newLeaseToken() (string, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func (s *JobStore) UpdateStatus(
	id int64,
	status model.JobStatus,
	targetStatus model.JobStatus,
	leaseToken string,
) error {
	result, err := s.db.Exec(
		`UPDATE jobs
         SET status = $1
         WHERE id = $2
           AND status = $3
           AND lease_token = $4
           AND lease_expires_at > NOW()`,
		targetStatus,
		id,
		status,
		leaseToken,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf(
			"job %d finish rejected: status, lease, or token no longer matches",
			id,
		)
	}

	return nil
}

func (s *JobStore) Finish(id int64, token string, status model.JobStatus, result []byte, errorMessage string) error {
	updated, err := s.db.Exec(
		`UPDATE jobs
		 SET status = $1, result = NULLIF($2::text, '')::jsonb, error_message = NULLIF($3::text, '')
		 WHERE id = $4 AND status = 'running'
		   AND lease_token = $5 AND lease_expires_at > NOW()`,
		status, string(result), errorMessage, id, token,
	)
	if err != nil {
		return err
	}
	affected, err := updated.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("job %d finish rejected: status, lease, or token no longer matches", id)
	}
	return nil
}

func (s *JobStore) ExtendLease(id int64, leaseToken string) error {
	result, err := s.db.Exec(
		`UPDATE jobs
		SET lease_expires_at = NOW() + (10 * INTERVAL '1 second')
		WHERE id = $1 AND status = 'running' AND lease_token = $2 AND lease_expires_at > NOW()`,
		id, leaseToken,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("job %d not found or not running", id)
	}

	return nil
}
