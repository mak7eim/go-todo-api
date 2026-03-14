# Task Manager API

A simple REST API for task management built with Go. This project demonstrates clean architecture patterns, separation of concerns, and best practices for building HTTP servers in Go.

## Features

- Create, read, update, and delete tasks
- In-memory storage (easily swappable for database)
- Clean architecture with separated concerns
- RESTful API design
- JSON request/response format

## Project Structure

todo-api
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── internal
│   ├── handlers
│   │   ├── task.go
│   │   └── tasks.go
│   ├── models
│   │   └── task.go
│   └── storage
│       └── memory.go
└── README.md