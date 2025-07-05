package repository

import (
	"sync"

	"go-gemini/domain"
)

// InMemoryTodoRepository implements domain.TodoRepository for in-memory storage.
type InMemoryTodoRepository struct {
	todos []domain.Todo
	mu    sync.Mutex // Mutex to protect concurrent access to todos
	nextID int
}

// NewInMemoryTodoRepository creates a new InMemoryTodoRepository.
func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos:  []domain.Todo{},
		nextID: 1,
	}
}

// Create adds a new Todo to the in-memory store.
func (r *InMemoryTodoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.ID = r.nextID
	r.nextID++
	r.todos = append(r.todos, *todo)
	return todo, nil
}
