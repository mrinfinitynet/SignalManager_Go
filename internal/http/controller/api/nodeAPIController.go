package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NodeAPIController struct{}

func NewNodeAPIController() *NodeAPIController {
	return &NodeAPIController{}
}

func (ctrl *NodeAPIController) Notification(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification received",
	})
}

func (ctrl *NodeAPIController) InitialTrades(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Initial trades received",
	})
}
