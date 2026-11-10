package biz

import (
	"backend-lab-go/internal/dal"
	"backend-lab-go/internal/model"
	"encoding/json"
	"log"
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
	return a.jobStore.ClaimNext()
}

func (a *JobBiz) Finish(id int64, runErr error) error {
	targetStatus := model.JobSucceeded

	if runErr != nil {
		targetStatus = model.JobFailed
	}

	job, err := a.jobStore.GetJobByID(id)
	if err != nil {
		log.Printf("Get job by id failed, id: %v", id)
		return err
	}

	if job.Status != model.JobRunning {
		log.Printf("Invalid job status: %v", job.Status)
		return nil
	}

	return a.jobStore.UpdateStatus(id, job.Status, targetStatus)
}
