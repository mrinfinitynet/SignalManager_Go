package routes

import (
	"SignalManager/internal/http/controller/api"
	"SignalManager/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		// node controller
		SignalFormatAPIController := api.NewSignalFormatAPIController()
		signalshotFormatGroup := apiGroup.Group("/signal-format").Use(middleware.APIKeyMiddleware())
		{
			signalshotFormatGroup.GET("/groups", SignalFormatAPIController.Groups)
			signalshotFormatGroup.POST("/submit-form", SignalFormatAPIController.SubmitForm)
		}

		// signal-shot controller
		signalShotController := api.NewSignalShotController()
		signalShotGroup := apiGroup.Group("/signal-shot").Use(middleware.APIKeyMiddleware())
		{
			signalShotGroup.GET("/active-real-positions", signalShotController.ActiveRealPositions)
			signalShotGroup.GET("/active-demo-positions", signalShotController.ActiveDemoPositions)
			signalShotGroup.POST("/risk-management-reset", signalShotController.RiskManagementReset)
			signalShotGroup.GET("/partial-templates", signalShotController.PartialTemplates)

			// set api keys
			signalShotGroup.POST("/bybit-api-keys", signalShotController.BybitApiKeys)
			signalShotGroup.POST("/binance-api-keys", signalShotController.BinanceApiKeys)
			signalShotGroup.POST("/get-api-keys", signalShotController.GetApiKeys)

			// crypto apis
			signalShotGroup.POST("/get-crypto-wallet", signalShotController.CryptoWallet)
			signalShotGroup.POST("/get-crypto-positions", signalShotController.CryptoPositions)
		}

		// node controller
		nodeController := api.NewNodeAPIController()
		nodeGroup := apiGroup.Group("/node").Use(middleware.APIKeyMiddleware())
		{
			nodeGroup.POST("/notification", nodeController.Notification)
			nodeGroup.POST("/ini-trades", nodeController.InitialTrades)
		}
	}
}
