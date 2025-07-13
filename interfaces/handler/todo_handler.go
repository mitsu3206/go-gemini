package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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
	Title string   `json:"title" binding:"required"`
	Tags  []string `json:"tags"`
}

// UpdateTodoRequest represents the request body for updating a Todo.
type UpdateTodoRequest struct {
	Title     string   `json:"title"`
	Completed bool     `json:"completed"`
	Tags      []string `json:"tags"`
}

// CreateTodo handles the creation of a new Todo item.
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.TodoUseCase.CreateTodo(req.Title, req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodo handles the retrieval of a single Todo item by ID.
func (h *TodoHandler) GetTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	todo, err := h.TodoUseCase.GetTodoByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, todo)
}

// GetAllTodos handles the retrieval of all Todo items.
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.TodoUseCase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// UpdateTodo handles the update of an existing Todo item.
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTodo, err := h.TodoUseCase.GetTodoByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	existingTodo.Title = req.Title
	existingTodo.Completed = req.Completed

	updatedTodo, err := h.TodoUseCase.UpdateTodo(existingTodo, req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

// DeleteTodo handles the deletion of a Todo item by ID.
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.TodoUseCase.DeleteTodo(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RemoveTagFromTodo handles the removal of a tag from a todo.
func (h *TodoHandler) RemoveTagFromTodo(c *gin.Context) {
	todoIDParam := c.Param("id")
	todoID, err := strconv.ParseUint(todoIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID format"})
		return
	}

	tagIDParam := c.Param("tag_id")
	tagID, err := strconv.ParseUint(tagIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Tag ID format"})
		return
	}

	err = h.TodoUseCase.RemoveTagFromTodo(uint(todoID), uint(tagID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
