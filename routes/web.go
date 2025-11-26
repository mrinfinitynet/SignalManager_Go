package routes

import (
	web "SignalManager/internal/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupWebRoutes(router *gin.Engine) {

	// web hooks
	SignalFormatAPIController := web.NewBotController()
	router.POST("/telegram-message-webhook", SignalFormatAPIController.Webhook)
}
