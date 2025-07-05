package usecase

import (
	"go-gemini/domain"
)

// TodoUseCase defines the use case for Todo operations.
type TodoUseCase struct {
	TodoRepo   domain.TodoRepository
	TagUseCase *TagUseCase
}

// NewTodoUseCase creates a new TodoUseCase.
func NewTodoUseCase(todoRepo domain.TodoRepository, tagUseCase *TagUseCase) *TodoUseCase {
	return &TodoUseCase{TodoRepo: todoRepo, TagUseCase: tagUseCase}
}

// CreateTodo creates a new Todo item.
func (uc *TodoUseCase) CreateTodo(title string, tagNames []string) (*domain.Todo, error) {
	todo := &domain.Todo{
		Title:     title,
		Completed: false,
	}

	if len(tagNames) > 0 {
		tags, err := uc.TagUseCase.GetOrCreateTags(tagNames)
		if err != nil {
			return nil, err
		}
		todo.Tags = tags
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

// UpdateTodo updates an existing Todo item.
func (uc *TodoUseCase) UpdateTodo(todo *domain.Todo, tagNames []string) (*domain.Todo, error) {
	if len(tagNames) > 0 {
		tags, err := uc.TagUseCase.GetOrCreateTags(tagNames)
		if err != nil {
			return nil, err
		}
		todo.Tags = tags
	} else {
		todo.Tags = []*domain.Tag{}
	}
	return uc.TodoRepo.Update(todo)
}

// DeleteTodo deletes a Todo item by its ID.
func (uc *TodoUseCase) DeleteTodo(id uint) error {
	return uc.TodoRepo.Delete(id)
}
