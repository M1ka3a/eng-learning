package dal

import (
	"backend-lab-go/internal/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

type JobStore struct {
	db *sql.DB
}

func NewJobStore(db *sql.DB) *JobStore {
	return &JobStore{db: db}
}

func (s *JobStore) Create(title string, ctx json.RawMessage) (model.Job, error) {
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
		`INSERT INTO jobs (title, ctx)
		 VALUES ($1, $2)
		 RETURNING id, title,status, ctx`,
		title, ctxJSON,
	).Scan(&job.ID, &job.Title, &job.Status, &job.Ctx)
	if err != nil {
		return model.Job{}, err
	}

	return job, nil
}

func (s *JobStore) GetJobByID(id int64) (model.Job, error) {
	var job model.Job

	err := s.db.QueryRow(
		`SELECT id, title, status, ctx
         FROM jobs
         WHERE id = $1`,
		id,
	).Scan(
		&job.ID,
		&job.Title,
		&job.Status,
		&job.Ctx,
	)

	if err != nil {
		return model.Job{}, err
	}

	return job, nil
}

func (s *JobStore) ClaimNext() (model.Job, error) {
	job := model.Job{}

	err := s.db.QueryRow(
		`UPDATE jobs
		SET status = 'running'
		WHERE id = (
			SELECT id
			FROM jobs
			WHERE status = 'pending'
			ORDER BY created_at, id
			FOR UPDATE SKIP LOCKED
			LIMIT 1
		)
		RETURNING id, title, status, ctx;`,
	).Scan(
		&job.ID,
		&job.Title,
		&job.Status,
		&job.Ctx,
	)
	if err != nil {
		return model.Job{}, err
	}
	return job, nil
}

func (s *JobStore) MarkSucceeded(id int64) error {
	result, err := s.db.Exec(
		`UPDATE jobs
		SET status = 'succeeded'
		WHERE id = $1 AND status = 'running'`,
		id,
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

func (s *JobStore) UpdateStatus(id int64, status model.JobStatus, targetStatus model.JobStatus) error {
	result, err := s.db.Exec(
		`UPDATE jobs
         SET status = $1
         WHERE id = $2 AND status = $3`,
		targetStatus,
		id,
		status,
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
			"job %d status transition failed: expected %s",
			id,
			status,
		)
	}

	return nil
}
