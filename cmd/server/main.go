package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sarfraz-4182/gin-pay/internal/api"
	"github.com/sarfraz-4182/gin-pay/internal/db"
	"github.com/sarfraz-4182/gin-pay/internal/middleware"
	"github.com/sarfraz-4182/gin-pay/internal/models"
)

func main() {
	_ = godotenv.Load() // loads .env

	db.Connect()

	// automigrate
	if err := db.DB.AutoMigrate(&models.Payment{}); err != nil {
		log.Fatalf("migrate error: %v", err)
	}

	r := gin.Default()

	// health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.ApiKeyAuth())
	api.RegisterPaymentRoutes(apiGroup)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
