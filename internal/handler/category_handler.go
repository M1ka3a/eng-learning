package handler

import (
	"backend-lab-go/internal/dal"
	"backend-lab-go/internal/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type CategoryHandler struct {
	store *dal.CategoryStore
}

func NewCategoryHandler(store *dal.CategoryStore) *CategoryHandler {
	return &CategoryHandler{store: store}
}

func (h *CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	categories, err := h.store.GetAll()
	if err != nil {
		http.Error(w, "Get categories failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	request := model.Category{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	newCategory, err := h.store.Add(request.Title)
	if err != nil {
		if errors.Is(err, model.ErrTitleInvalid) {
			http.Error(w, "invalid title", http.StatusBadGateway)
			return
		}
		log.Printf("Create category failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newCategory); err != nil {
		log.Println("Encode response failed:", err)
	}
}
