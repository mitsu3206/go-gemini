package usecase

import (
	"go-gemini/domain"
)

// TodoUseCase defines the use case for Todo operations.
type TodoUseCase struct {
	TodoRepo domain.TodoRepository
}

// NewTodoUseCase creates a new TodoUseCase.
func NewTodoUseCase(repo domain.TodoRepository) *TodoUseCase {
	return &TodoUseCase{TodoRepo: repo}
}

// CreateTodo creates a new Todo item.
func (uc *TodoUseCase) CreateTodo(title string) (*domain.Todo, error) {
	todo := &domain.Todo{
		Title:     title,
		Completed: false,
	}
	return uc.TodoRepo.Create(todo)
}