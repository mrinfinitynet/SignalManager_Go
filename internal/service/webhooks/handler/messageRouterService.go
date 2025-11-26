package handler

import (
	// "SignalManager/pkg/logger"
	"github.com/gin-gonic/gin"
)

type MessageRouterService struct{}

func NewMessageRouterService() *MessageRouterService {
	return &MessageRouterService{}
}

func (s *MessageRouterService) Handle(c *gin.Context, data map[string]any) {
	// chat_id, _            := message["chat_id"].(map[string]any)
	// forward_id, _            := message["forward_id"].(map[string]any)
	// text, _            := message["text"].(map[string]any)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
