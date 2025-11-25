package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
	AI       *AIConfig
	Mail     *MailConfig
}

var GlobalConfig *Config

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	GlobalConfig = &Config{
		Server:   loadServerConfig(),
		Database: loadDatabaseConfig(),
		AI:       loadAIConfig(),
		Mail:     loadMailConfig(),
	}

	log.Printf("Config loaded successfully")
	log.Printf("Server Port: %s", GlobalConfig.Server.Port)
	log.Printf("Database: %s@%s:%s/%s", GlobalConfig.Database.User, GlobalConfig.Database.Host, GlobalConfig.Database.Port, GlobalConfig.Database.DBName)
	log.Printf("AI Provider: %s", GlobalConfig.AI.Provider)
	log.Printf("Mail From: %s <%s>", GlobalConfig.Mail.FromName, GlobalConfig.Mail.FromEmail)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
