package domain

import "time"

// Todo represents a single todo item.
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoRepository defines the interface for todo data storage.
type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
}