package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockyprabowo/assignment-2/models"
)

func (controller OrderController) Create(context *gin.Context) {
	var (
		order models.Order
		err   error
	)
	if err := context.ShouldBind(&order); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if err = controller.Actions.CreateOrder(&order); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": order})
}
