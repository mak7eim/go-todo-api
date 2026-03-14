package main

import (
	"log"
	"net/http"
	"todo/internal/handlers"
	"todo/internal/storage"
)

func main() {
	store := storage.NewStorageMemory()

	taskHandler := handlers.NewTaskHandler(store)
	tasksHandler := handlers.NewTasksHandler(store)

	http.Handle("/tasks", tasksHandler)
	http.Handle("/tasks/", taskHandler)

	log.Println("server start")

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
