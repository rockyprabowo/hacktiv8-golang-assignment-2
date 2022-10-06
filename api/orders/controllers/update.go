package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockyprabowo/assignment-2/models"
)

func (controller OrderController) Update(context *gin.Context) {
	var (
		id          = context.Param("id")
		order       models.Order
		orderUpdate models.Order
		err         error
	)

	if err := context.ShouldBind(&orderUpdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	order, err = controller.Actions.GetOrderById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	err = controller.Actions.UpdateOrder(&order, &orderUpdate)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": order, "message": "Updated"})
}
