package worker

import (
	"backend-lab-go/internal/biz"
	"backend-lab-go/internal/model"
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

type Worker struct {
	jobBiz *biz.JobBiz
}

func NewWorker(jobBiz *biz.JobBiz) *Worker {
	return &Worker{jobBiz: jobBiz}
}

func (w *Worker) Execute(job model.Job) error {
	// TODO: 后续替换成真正的任务执行逻辑
	time.Sleep(5 * time.Second)
	return nil
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	log.Println("Worker started")

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker stopped")
			return

		case <-ticker.C:
			job, err := w.jobBiz.ClaimNext()

			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					continue
				}

				log.Printf("Claim job failed: %v", err)
				continue
			}

			log.Printf("Claim job successfully, id=%d", job.ID)

			// 模拟执行任务
			runErr := w.Execute(job)

			if err := w.jobBiz.Finish(job.ID, job.LeaseToken, runErr); err != nil {
				log.Printf(
					"Mark job %d succeeded failed: %v",
					job.ID,
					err,
				)
				continue
			}

			log.Printf("Job %d succeeded", job.ID)
		}
	}
}
