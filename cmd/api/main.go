package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var currentID = 1

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)

	log.Println("server start")

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getTasks(w)
	case http.MethodPost:
		postTasks(w, r)
	default:
		http.Error(w, "не верный метод", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(tasks)

	if err != nil {
		http.Error(w, "ошибка кодирования", http.StatusInternalServerError)
	}
}

func postTasks(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "неверный формат JSON", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(w, "поле title обязательно", http.StatusBadRequest)
		return
	}

	task.ID = strconv.Itoa(currentID)
	currentID++

	tasks = append(tasks, task)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		log.Println("ошибка кодировки", http.StatusInternalServerError)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")

	id := parts[2]

	if id != "" {
		http.Error(w, "id не указан", http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodGet:
		getTask(w, id)
	case http.MethodPut:
		putTask(w, r, id)
	case http.MethodDelete:
		deleteTask(w, id)
	default:
		http.Error(w, "неверный метод", http.StatusMethodNotAllowed)
	}
}

func getTask(w http.ResponseWriter, id string) {
	for _, v := range tasks {
		if v.ID == id {
			err := json.NewEncoder(w).Encode(v)

			if err != nil {
				http.Error(w, "ошибка кодировки", http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "задача не найдена", http.StatusNotFound)
}

func putTask(w http.ResponseWriter, r *http.Request, id string) {
	var updatedTask Task

	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "неверный формат JSON", http.StatusBadRequest)
		return
	}

	if updatedTask.Title == "" {
		http.Error(w, "поле title обязательно", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	http.Error(w, "задача не найдена", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, id string) {
	for i, v := range tasks {
		if v.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "задача не найдена", http.StatusNotFound)
}
