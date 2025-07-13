package repository

import (
	"go-gemini/domain"

	"gorm.io/gorm"
)

// GormTodoRepository implements domain.TodoRepository using GORM.
type GormTodoRepository struct {
	DB *gorm.DB
}

// NewGormTodoRepository creates a new GormTodoRepository.
func NewGormTodoRepository(db *gorm.DB) *GormTodoRepository {
	return &GormTodoRepository{DB: db}
}

// Create adds a new Todo to the database using GORM.
func (r *GormTodoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	result := r.DB.Create(todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

// FindByID retrieves a Todo item by its ID using GORM.
func (r *GormTodoRepository) FindByID(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	result := r.DB.Preload("Tags").First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// FindAll retrieves all Todo items using GORM.
func (r *GormTodoRepository) FindAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	result := r.DB.Preload("Tags").Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

// Update updates an existing Todo item in the database using GORM.
func (r *GormTodoRepository) Update(todo *domain.Todo) (*domain.Todo, error) {
	result := r.DB.Save(todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

// Delete deletes a Todo item by its ID using GORM.
func (r *GormTodoRepository) Delete(id uint) error {
	result := r.DB.Delete(&domain.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// RemoveTagFromTodo removes a tag from a todo.
func (r *GormTodoRepository) RemoveTagFromTodo(todoID, tagID uint) error {
	return r.DB.Where("todo_id = ? AND tag_id = ?", todoID, tagID).Delete(&domain.TodoTag{}).Error
}
