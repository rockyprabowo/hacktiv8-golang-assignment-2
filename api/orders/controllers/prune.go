package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller OrderController) Prune(context *gin.Context) {
	count, err := controller.Actions.Prune()
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Prune failed."},
		)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"count":   count,
		"message": "Prune complete.",
	})
}
