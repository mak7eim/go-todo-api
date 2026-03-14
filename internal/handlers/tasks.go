package handlers

import (
	"encoding/json"
	"net/http"
	"todo/internal/models"
	"todo/internal/storage"
)

type TasksHandler struct {
	storage *storage.MemoryStorage
}

func NewTasksHandler(s *storage.MemoryStorage) *TasksHandler {
	return &TasksHandler{
		storage: s,
	}
}

func (h *TasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.getTasks(w)
	case http.MethodPost:
		h.postTasks(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TasksHandler) getTasks(w http.ResponseWriter) {
	tasks := h.storage.GetAll()
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "failed to encode tasks", http.StatusInternalServerError)
	}
}

func (h *TasksHandler) postTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid JSON format", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	created := h.storage.Create(task)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(created); err != nil {
		http.Error(w, "failed to encode task", http.StatusInternalServerError)
	}
}
