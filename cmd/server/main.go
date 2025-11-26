package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"SignalManager/pkg/config"
	"SignalManager/pkg/logger"
	"SignalManager/routes"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize logger
	if err := logger.Initialize("storage/logs"); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Close()

	logger.Info("Logger initialized successfully")

	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	logger.Info("Server starting on :%s", config.GlobalConfig.Server.Port)
	if err := router.Run(":" + config.GlobalConfig.Server.Port); err != nil {
		logger.Error("Failed to start server: %v", err)
		log.Fatal("Failed to start server:", err)
	}
}
