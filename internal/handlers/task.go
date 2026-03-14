package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo/internal/models"
	"todo/internal/storage"
)

type TaskHandler struct {
	storage *storage.MemoryStorage
}

func NewTaskHandler(s *storage.MemoryStorage) *TaskHandler {
	return &TaskHandler{
		storage: s,
	}
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := h.extractID(r.URL.Path)
	if id == "" {
		http.Error(w, "id not specified", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getTask(w, id)
	case http.MethodPut:
		h.putTask(w, r, id)
	case http.MethodDelete:
		h.deleteTask(w, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) extractID(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

func (h *TaskHandler) getTask(w http.ResponseWriter, id string) {
	task, found := h.storage.GetByID(id)
	if !found {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "failed to encode task", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) putTask(w http.ResponseWriter, r *http.Request, id string) {
	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "invalid JSON format", http.StatusBadRequest)
		return
	}

	if updatedTask.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	task, found := h.storage.Update(id, updatedTask)
	if !found {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "failed to encode task", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, id string) {
	if h.storage.Delete(id) {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.Error(w, "task not found", http.StatusNotFound)
}
