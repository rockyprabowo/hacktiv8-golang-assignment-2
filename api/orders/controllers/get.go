package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockyprabowo/assignment-2/api/orders/responses"
	"rockyprabowo/assignment-2/models"
)

// GetById
// @Summary     Fetch single order
// @Description Fetch an order by its ID from the database.
// @Tags        orders
// @Produce     json
// @Param       id  path     int true "Order ID"
// @Success     200 {object} order_responses.Data
// @Failure     404 {object} order_responses.Error
// @Failure     500 {object} order_responses.Error
// @Router      /orders/{id} [get]
func (controller OrderController) GetById(context *gin.Context) {
	var (
		id    = context.Param("id")
		order models.Order
		err   error
	)

	order, err = controller.Actions.GetOrderById(id)
	if err != nil {
		context.JSON(
			http.StatusNotFound,
			order_responses.Error{
				Message: err.Error(),
				Status:  "error",
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		order_responses.Data{
			Data:    order,
			Message: "found",
		},
	)
}
