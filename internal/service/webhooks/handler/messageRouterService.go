package handler

import (
	"SignalManager/pkg/logger"
	// "github.com/gin-gonic/gin"
)

type MessageRouterService struct{}

func NewMessageRouterService() *MessageRouterService {
	return &MessageRouterService{}
}

func (s *MessageRouterService) Handle(c *gin.Context, message map[string]any) {
	// Get data
	// chat, _ := message["chat"].(map[string]any)
	// chatId, _ := chat["id"].(float64)
	// logger.Info("Chat Data: %v Chat ID: %v", chat, chatId)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
