package domain

import (
	"gorm.io/gorm"
)

// Todo represents a single todo item.
type Todo struct {
	gorm.Model
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}

// TodoRepository defines the interface for todo data storage.
type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
	FindByID(id uint) (*Todo, error)
	FindAll() ([]*Todo, error)
}
