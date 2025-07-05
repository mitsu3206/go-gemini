package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gemini/usecase"
)

// TodoHandler handles HTTP requests related to Todo items.
type TodoHandler struct {
	TodoUseCase *usecase.TodoUseCase
}

// NewTodoHandler creates a new TodoHandler.
func NewTodoHandler(uc *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{TodoUseCase: uc}
}

// CreateTodoRequest represents the request body for creating a Todo.
type CreateTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

// CreateTodo handles the creation of a new Todo item.
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.TodoUseCase.CreateTodo(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
