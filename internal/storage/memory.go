package storage

import (
	"strconv"
	"todo/internal/models"
)

type MemoryStorage struct {
	tasks     []models.Task
	currentID int
}

func NewStorageMemory() *MemoryStorage {
	return &MemoryStorage{
		tasks:     []models.Task{},
		currentID: 1,
	}
}

func (s *MemoryStorage) GetAll() []models.Task {
	return s.tasks
}

func (s *MemoryStorage) GetByID(id string) (models.Task, bool) {
	for _, v := range s.tasks {
		if v.ID == id {
			return v, true
		}
	}
	return models.Task{}, false
}

func (s *MemoryStorage) Create(task models.Task) models.Task {
	task.ID = strconv.Itoa(s.currentID)
	s.currentID++
	s.tasks = append(s.tasks, task)
	return task
}

func (s *MemoryStorage) Update(id string, updatedTask models.Task) (models.Task, bool) {
	for i, task := range s.tasks {
		if task.ID == id {
			updatedTask.ID = id
			s.tasks[i] = updatedTask
			return updatedTask, true
		}
	}
	return models.Task{}, false
}

func (s *MemoryStorage) Delete(id string) bool {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}
