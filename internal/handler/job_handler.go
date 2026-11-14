package handler

import (
	"backend-lab-go/internal/biz"
	"backend-lab-go/internal/model"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type JobHandler struct {
	jobBiz *biz.JobBiz
}

func NewJobHandler(jobBiz *biz.JobBiz) *JobHandler {
	return &JobHandler{jobBiz: jobBiz}
}

type CreateJobRequest struct {
	Title string          `json:"title"`
	Ctx   json.RawMessage `json:"ctx"`
}

type SubmitOrderIDRequest struct {
	UID     string `json:"uid"`
	OrderID string `json:"order_id"`
}

// POST /jobs
func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateJobRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	job, err := h.jobBiz.Create(request.Title, request.Ctx)
	if err != nil {
		if errors.Is(err, model.ErrJobTitleInvalid) ||
			errors.Is(err, model.ErrJobCtxInvalid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Create job failed: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(job); err != nil {
		log.Printf("Encode response failed: %v", err)
	}
}

// PATCH /jobs/{id}/input
func (h *JobHandler) SubmitOrderID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid job ID", http.StatusBadRequest)
		return
	}

	var request SubmitOrderIDRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	job, err := h.jobBiz.AddOrderID(id, request.UID, request.OrderID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "job not found", http.StatusNotFound)
		case errors.Is(err, model.ErrJobUIDInvalid), errors.Is(err, model.ErrJobOrderInvalid):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, model.ErrJobUIDMismatch):
			http.Error(w, err.Error(), http.StatusForbidden)
		case errors.Is(err, model.ErrJobNotEditable):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			log.Printf("Update job failed: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(job); err != nil {
		log.Printf("Encode response failed: %v", err)
	}
}

// GET /jobs/{id}
func (h *JobHandler) GetJobByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid job ID", http.StatusBadRequest)
		return
	}

	job, err := h.jobBiz.GetJobByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "job not found", http.StatusNotFound)
			return
		}

		log.Printf("Get job failed: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(job); err != nil {
		log.Printf("Encode response failed: %v", err)
	}
}
