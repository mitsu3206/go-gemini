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

// GetTodoByID retrieves a Todo item by its ID.
func (uc *TodoUseCase) GetTodoByID(id uint) (*domain.Todo, error) {
	return uc.TodoRepo.FindByID(id)
}

// GetAllTodos retrieves all Todo items.
func (uc *TodoUseCase) GetAllTodos() ([]*domain.Todo, error) {
	return uc.TodoRepo.FindAll()
}