package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockyprabowo/assignment-2/api/orders/responses"
	"rockyprabowo/assignment-2/models"
)

// GetAll godoc
// @Summary     Fetch all orders
// @Description Fetch all order(s) from the database.
// @Tags        orders
// @Produce     json
// @Success     200 {object} order_responses.DataList
// @Failure     500 {object} order_responses.Error
// @Router      /orders [get]
func (controller OrderController) GetAll(context *gin.Context) {
	var (
		orders []models.Order
		err    error
	)

	if orders, err = controller.Actions.GetAllOrders(); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			order_responses.Error{
				Message: err.Error(),
				Status:  "error",
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		order_responses.DataList{
			Data:    orders,
			Count:   len(orders),
			Message: "success",
		},
	)
}
