package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"SignalManager/pkg/config"
	"SignalManager/routes"
)

func main() {
	// Load configuration
	config.LoadConfig()

	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	log.Printf("Server starting on :%s", config.GlobalConfig.Server.Port)
	if err := router.Run(":" + config.GlobalConfig.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
