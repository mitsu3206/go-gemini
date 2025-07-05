package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"go-gemini/domain"
	"go-gemini/infrastructure/repository"
	"go-gemini/interfaces/handler"
	"go-gemini/usecase"
)

func main() {
	// Database connection with retry
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),	
		os.Getenv("DB_PORT"),
	)

	var db *gorm.DB
	var err error
	maxRetries := 10
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(driver.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			log.Println("Successfully connected to database.")
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v. Retrying in %s...", i+1, maxRetries, err, retryInterval)
		time.Sleep(retryInterval)
	}

	if err != nil {
		log.Fatalf("Exceeded max retries to connect to database: %v", err)
	}

	// AutoMigrate will create/update table based on struct
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	r := gin.Default()

	// Dependencies Injection
	todoRepo := repository.NewGormTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUseCase)

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos/:id", todoHandler.GetTodo)
	r.GET("/todos", todoHandler.GetAllTodos)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}