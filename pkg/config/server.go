package config

type ServerConfig struct {
	Port        string
	Environment string
	Debug       bool
	AllowOrigin string
	JWTToken    string
}

func loadServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:        getEnv("APP_PORT", "8080"),
		Environment: getEnv("APP_ENV", "development"),
		Debug:       getEnv("APP_DEBUG", "true") == "true",
		AllowOrigin: getEnv("ALLOW_ORIGIN", "*"),
		JWTToken:    getEnv("JWT_TOKEN", ""),
	}
}
