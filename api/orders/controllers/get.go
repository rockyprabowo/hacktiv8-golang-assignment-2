package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockyprabowo/assignment-2/models"
)

func (controller OrderController) GetById(context *gin.Context) {
	var (
		id    = context.Param("id")
		order models.Order
		err   error
	)

	order, err = controller.Actions.GetOrderById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": order, "count": 1})
}
