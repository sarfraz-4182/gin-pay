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
		// Get memory info
		// memInfo, _ := mem.VirtualMemory()

		// Get CPU usage (per second)
		// cpuPercent, _ := cpu.Percent(time.Second, false)

		// Get disk usage of root ("/")
		// diskInfo, _ := disk.Usage("/")

		c.JSON(200, gin.H{
			"status": "ok",
			// "system": gin.H{
			// 	"cpu_usage":     cpuPercent[0],              // percentage
			// 	"memory_used":   memInfo.Used / 1024 / 1024, // in MB
			// 	"memory_total":  memInfo.Total / 1024 / 1024,
			// 	"memory_used_%": memInfo.UsedPercent,
			// 	"disk_used":     diskInfo.Used / 1024 / 1024 / 1024, // in GB
			// 	"disk_total":    diskInfo.Total / 1024 / 1024 / 1024,
			// 	"disk_used_%":   diskInfo.UsedPercent,
			// },
		})
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
