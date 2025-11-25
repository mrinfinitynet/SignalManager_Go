package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BotController struct{}

func NewBotController() *BotController {
	return &BotController{}
}

func (ctrl *BotController) Webhook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Webhook received",
	})
}
