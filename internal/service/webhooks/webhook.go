package webhooks

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type WebhookService struct{}

func NewWebhookService() *WebhookService {
	return &WebhookService{}
}

func (s *WebhookService) HandleWebhook(c *gin.Context) error {
	fmt.Println("webhook")
	return nil
}
