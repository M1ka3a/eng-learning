package biz

import (
	"backend-lab-go/internal/dal"
	"backend-lab-go/internal/model"
	"encoding/json"
	"strings"
)

type JobBiz struct {
	jobStore *dal.JobStore
}

type TicketDetails struct {
	UID         string               `json:"uid"`
	OrderID     string               `json:"order_id"`
	Description string               `json:"description"`
	Category    model.TicketCategory `json:"category"`
}

func NewJobBiz(jobStore *dal.JobStore) *JobBiz {
	return &JobBiz{jobStore: jobStore}
}

func (a *JobBiz) Create(title string, ctx json.RawMessage) (model.Job, error) {
	var details TicketDetails
	if err := json.Unmarshal(ctx, &details); err != nil {
		return model.Job{}, model.ErrJobCtxInvalid
	}
	if details.UID == "" || details.UID != strings.TrimSpace(details.UID) {
		return model.Job{}, model.ErrJobUIDInvalid
	}
	if strings.TrimSpace(details.Description) == "" {
		return model.Job{}, model.ErrJobDescriptionInvalid
	}
	if !validTicketCategory(details.Category) {
		return model.Job{}, model.ErrJobCategoryInvalid
	}

	if strings.TrimSpace(details.OrderID) == "" {
		return a.jobStore.CreateWithStatus(title, model.JobNeedsInput, ctx)
	}
	return a.jobStore.CreateWithStatus(title, model.JobPending, ctx)
}

func (a *JobBiz) RouteTicket(ctx json.RawMessage) (model.RoutingResult, error) {
	var details TicketDetails
	if err := json.Unmarshal(ctx, &details); err != nil {
		return model.RoutingResult{}, model.ErrJobCtxInvalid
	}
	if strings.TrimSpace(details.UID) == "" || strings.TrimSpace(details.Description) == "" || strings.TrimSpace(details.OrderID) == "" {
		return model.RoutingResult{}, model.ErrJobCtxInvalid
	}
	if !validTicketCategory(details.Category) {
		return model.RoutingResult{}, model.ErrJobCategoryInvalid
	}
	result := model.RoutingResult{Category: details.Category, CategorySource: "user_selected"}
	switch details.Category {
	case model.TicketPaymentIssue:
		result.Queue = "payment_support"
		result.NextChecks = []string{"核对支付记录", "核对订单支付状态"}
	case model.TicketOrderIssue:
		result.Queue = "order_support"
		result.NextChecks = []string{"核对订单详情"}
	case model.TicketOther:
		result.Queue = "general_support"
		result.NextChecks = []string{"阅读用户描述并确认具体诉求"}
	}
	return result, nil
}

func validTicketCategory(category model.TicketCategory) bool {
	switch category {
	case model.TicketPaymentIssue, model.TicketOrderIssue, model.TicketOther:
		return true
	default:
		return false
	}
}

func (a *JobBiz) GetJobByID(id int64) (model.Job, error) {
	return a.jobStore.GetJobByID(id)
}

func (a *JobBiz) ListByUID(uid string) ([]model.Job, error) {
	uid = strings.TrimSpace(uid)
	if uid == "" {
		return nil, model.ErrJobUIDInvalid
	}
	return a.jobStore.ListByUID(uid)
}

func (a *JobBiz) ClaimNext() (model.Job, error) {
	return a.jobStore.ClaimNext(60)
}

func (a *JobBiz) Finish(id int64, token string, result model.RoutingResult, runErr error) error {
	targetStatus := model.JobSucceeded
	var resultJSON []byte
	var errorMessage string

	if runErr != nil {
		targetStatus = model.JobFailed
		errorMessage = runErr.Error()
	} else {
		var err error
		resultJSON, err = json.Marshal(result)
		if err != nil {
			return err
		}
	}

	return a.jobStore.Finish(id, token, targetStatus, resultJSON, errorMessage)
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
