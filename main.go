package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gemini/infrastructure/repository"
	"go-gemini/interfaces/handler"
	"go-gemini/usecase"
)

func main() {
	r := gin.Default()

	// Dependencies Injection
	todoRepo := repository.NewInMemoryTodoRepository()
	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUseCase)

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/todos", todoHandler.CreateTodo)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
