package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockyprabowo/assignment-2/models"
)

func (controller OrderController) GetAll(context *gin.Context) {
	var (
		orders []models.Order
		err    error
	)

	if orders, err = controller.Actions.GetAllOrders(); err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": err.Error(), "status": "error"},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"data": orders, "count": len(orders)},
	)
}
