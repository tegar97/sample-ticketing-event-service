package main

import (
	"log"
	"net/http"
	"os"

	"event-service/internal/config"
	"event-service/internal/handler"
	"event-service/internal/models"
	"event-service/internal/repository"
	"event-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	cfg := config.Load()

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.Event{}); err != nil {
		log.Fatal("Failed to auto migrate database:", err)
	}

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Event Service API",
			"version": "v1",
			"status":  "running",
		})
	})

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "event-service",
		})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/events", eventHandler.GetEvents)
		api.GET("/events/:id", eventHandler.GetEvent)
		api.POST("/events", eventHandler.CreateEvent)
		api.PUT("/events/:id", eventHandler.UpdateEvent)
		api.DELETE("/events/:id", eventHandler.DeleteEvent)
		api.PUT("/events/:id/tickets", eventHandler.UpdateAvailableTickets)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	log.Printf("Event service starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
