package biz

import (
	"backend-lab-go/internal/dal"
	"backend-lab-go/internal/model"
	"encoding/json"
	"fmt"
	"strings"
)

type JobBiz struct {
	jobStore *dal.JobStore
}

type TicketDetails struct {
	UID         string `json:"uid"`
	OrderID     string `json:"order_id"`
	Description string `json:"description"`
}

func NewJobBiz(jobStore *dal.JobStore) *JobBiz {
	return &JobBiz{jobStore: jobStore}
}

func (a *JobBiz) Create(title string, ctx json.RawMessage) (model.Job, error) {
	details := TicketDetails{}

	if err := json.Unmarshal(ctx, &details); err != nil {
		return model.Job{}, err
	}

	if strings.TrimSpace(details.UID) == "" {
		return model.Job{}, fmt.Errorf("Invalid uid")}
	}

	if strings.TrimSpace(details.OrderID) == "" {
		return a.jobStore.Create(title, model.JobNeedsInput, ctx)
	}
	return a.jobStore.Create(title, model.JobPending, ctx)
}

func (a *JobBiz) GetJobByID(id int64) (model.Job, error) {
	return a.jobStore.GetJobByID(id)
}

func (a *JobBiz) ClaimNext() (model.Job, error) {
	return a.jobStore.ClaimNext(60)
}

func (a *JobBiz) Finish(id int64, token string, runErr error) error {
	targetStatus := model.JobSucceeded

	if runErr != nil {
		targetStatus = model.JobFailed
	}

	return a.jobStore.UpdateStatus(id, model.JobRunning, targetStatus, token)
}

func (a *JobBiz) AddOrderID(id int64, uid, orderID string) (model.Job, error) {
	uid = strings.TrimSpace(uid)
	orderID = strings.TrimSpace(orderID)
	if uid == "" {
		return model.Job{}, model.ErrJobUIDInvalid
	}
	if orderID == "" {
		return model.Job{}, model.ErrJobOrderInvalid
	}

	job, err := a.jobStore.GetJobByID(id)
	if err != nil {
		return model.Job{}, err
	}
	var details TicketDetails
	if err := json.Unmarshal(job.Ctx, &details); err != nil {
		return model.Job{}, err
	}
	if details.UID != uid {
		return model.Job{}, model.ErrJobUIDMismatch
	}
	if job.Status != model.JobNeedsInput {
		return model.Job{}, model.ErrJobNotEditable
	}

	return a.jobStore.AddOrderID(id, uid, orderID)
}
