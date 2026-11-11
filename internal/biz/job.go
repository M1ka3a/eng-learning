package biz

import (
	"backend-lab-go/internal/dal"
	"backend-lab-go/internal/model"
	"encoding/json"
)

type JobBiz struct {
	jobStore *dal.JobStore
}

func NewJobBiz(jobStore *dal.JobStore) *JobBiz {
	return &JobBiz{jobStore: jobStore}
}

func (a *JobBiz) Create(title string, ctx json.RawMessage) (model.Job, error) {
	return a.jobStore.Create(title, ctx)
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
