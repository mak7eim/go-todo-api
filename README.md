# Todo API (Go)

A simple REST API for managing tasks written in **Go** using the standard **net/http** package.
The project demonstrates basic backend concepts such as routing, handlers, in-memory storage, and CRUD operations.
---

## Features

* Create a task
* Get all tasks
* Get task by ID
* Update task
* Delete task
* JSON request/response format
* In-memory storage
* Clean project structure using `internal` packages

---

## Project Structure

```
todo-api
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── handlers
│   │   ├── task.go
│   │   └── tasks.go
│   ├── models
│   │   └── task.go
│   └── storage
│       └── memory.go
├── go.mod
└── README.md
```

---

## Task Model

```
{
  "id": "1",
  "title": "Learn Go",
  "completed": false
}
```

---

## API Endpoints

### Get all tasks

```
GET /tasks
```

Response:

```
[
  {
    "id": "1",
    "title": "Learn Go",
    "completed": false
  }
]
```

---

### Create task

```
POST /tasks
```

Request body:

```
{
  "title": "Build REST API",
  "completed": false
}
```

---

### Get task by ID

```
GET /tasks/{id}
```

---

### Update task

```
PUT /tasks/{id}
```

Request body:

```
{
  "title": "Learn Go deeply",
  "completed": true
}
```

---

### Delete task

```
DELETE /tasks/{id}
```

---

## Running the Project

### 1. Clone the repository

```
git clone https://github.com/mak7eim/fo-todo-api.git
cd todo-api
```

### 2. Run the server

```
go run ./cmd/api
```

Server will start on:

```
http://localhost:8080
```

---

## Technologies

* Go
* net/http
* JSON

---