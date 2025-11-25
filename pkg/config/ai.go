package config

type AIConfig struct {
	OpenAIAPIKey    string
	OpenAIModel     string
	AnthropicAPIKey string
	AnthropicModel  string
	Provider        string // "openai" or "anthropic"
}

func loadAIConfig() *AIConfig {
	return &AIConfig{
		OpenAIAPIKey:    getEnv("OPENAI_API_KEY", "sk-demo-key-xxxxxxxxxxxxx"),
		OpenAIModel:     getEnv("OPENAI_MODEL", "gpt-4-turbo-preview"),
		AnthropicAPIKey: getEnv("ANTHROPIC_API_KEY", "sk-ant-demo-key-xxxxxxxxxxxxx"),
		AnthropicModel:  getEnv("ANTHROPIC_MODEL", "claude-3-5-sonnet-20241022"),
		Provider:        getEnv("AI_PROVIDER", "openai"),
	}
}
