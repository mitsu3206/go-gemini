package domain

import (
	"gorm.io/gorm"
)

// Todo represents a single todo item.
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Tags      []*Tag `json:"tags" gorm:"many2many:todo_tags;"`
}

// TodoRepository defines the interface for todo data storage.
type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
	FindByID(id uint) (*Todo, error)
	FindAll() ([]*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id uint) error
}
