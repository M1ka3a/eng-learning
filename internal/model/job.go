package model

import "encoding/json"

type Job struct {
	ID             int64           `json:"id"`
	Title          string          `json:"title"`
	Status         JobStatus       `json:"status"`
	Ctx            json.RawMessage `json:"ctx"`
	Result         json.RawMessage `json:"result,omitempty"`
	ErrorMessage   string          `json:"error,omitempty"`
	LeaseToken     string          `json:"-"`
	LeaseExpiresAt string          `json:"-"`
}

type JobStatus string

const (
	JobPending    JobStatus = "pending"
	JobRunning    JobStatus = "running"
	JobSucceeded  JobStatus = "succeeded"
	JobFailed     JobStatus = "failed"
	JobNeedsInput JobStatus = "needs_input"
)

type TicketCategory string

const (
	TicketPaymentIssue TicketCategory = "payment_issue"
	TicketOrderIssue   TicketCategory = "order_issue"
	TicketOther        TicketCategory = "other"
)

type RoutingResult struct {
	Category       TicketCategory `json:"category"`
	CategorySource string         `json:"category_source"`
	Queue          string         `json:"queue"`
	NextChecks     []string       `json:"next_checks"`
}
