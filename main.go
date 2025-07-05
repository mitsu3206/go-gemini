package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	driver "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
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
	err = db.AutoMigrate(&domain.Todo{}, &domain.Tag{})
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	r := gin.Default()

	todoRepo := repository.NewGormTodoRepository(db)
	tagRepo := repository.NewGormTagRepository(db)
	tagUseCase := usecase.NewTagUseCase(tagRepo)
	todoUseCase := usecase.NewTodoUseCase(todoRepo, tagUseCase)
	todoHandler := handler.NewTodoHandler(todoUseCase)
	tagHandler := handler.NewTagHandler(tagUseCase)

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
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// Tag Routes
	r.POST("/tags", tagHandler.CreateTag)
	r.GET("/tags/:id", tagHandler.GetTag)
	r.GET("/tags", tagHandler.GetAllTags)
	r.PUT("/tags/:id", tagHandler.UpdateTag)
	r.DELETE("/tags/:id", tagHandler.DeleteTag)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
