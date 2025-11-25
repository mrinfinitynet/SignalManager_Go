package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignalFormatAPIController struct{}

func NewSignalFormatAPIController() *SignalFormatAPIController {
	return &SignalFormatAPIController{}
}

func (ctrl *SignalFormatAPIController) Groups(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Groups received",
	})
}

func (ctrl *SignalFormatAPIController) SubmitForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SubmitForm received",
	})
}
