package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignalShotAPIController struct{}

func NewSignalShotController() *SignalShotAPIController {
	return &SignalShotAPIController{}
}

func (ctrl *SignalShotAPIController) ActiveRealPositions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "activeRealPositions received",
	})
}

func (ctrl *SignalShotAPIController) ActiveDemoPositions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "activeDemoPositions received",
	})
}

func (ctrl *SignalShotAPIController) RiskManagementReset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "riskManagementReset received",
	})
}

func (ctrl *SignalShotAPIController) PartialTemplates(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "partialTemplates received",
	})
}

func (ctrl *SignalShotAPIController) BybitApiKeys(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "bybitApiKeys received",
	})
}

func (ctrl *SignalShotAPIController) BinanceApiKeys(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "binanceApiKeys received",
	})
}

func (ctrl *SignalShotAPIController) GetApiKeys(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getApiKeys received",
	})
}

func (ctrl *SignalShotAPIController) CryptoWallet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "cryptoWallet received",
	})
}

func (ctrl *SignalShotAPIController) CryptoPositions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "cryptoPositions received",
	})
}
