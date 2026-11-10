package model

import "encoding/json"

type Job struct {
	ID     int64           `json:"id"`
	Title  string          `json:"title"`
	Status JobStatus       `json:"status"`
	Ctx    json.RawMessage `json:"ctx"`
}

type JobStatus string

const (
	JobPending   JobStatus = "pending"
	JobRunning   JobStatus = "running"
	JobSucceeded JobStatus = "succeeded"
	JobFailed    JobStatus = "failed"
)
