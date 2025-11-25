package config

type MailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	FromEmail    string
	FromName     string
}

func loadMailConfig() *MailConfig {
	return &MailConfig{
		SMTPHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:     getEnv("SMTP_PORT", "587"),
		SMTPUser:     getEnv("SMTP_USER", "demo@example.com"),
		SMTPPassword: getEnv("SMTP_PASSWORD", "demo-password"),
		FromEmail:    getEnv("MAIL_FROM_EMAIL", "noreply@signalmanager.com"),
		FromName:     getEnv("MAIL_FROM_NAME", "Signal Manager"),
	}
}
