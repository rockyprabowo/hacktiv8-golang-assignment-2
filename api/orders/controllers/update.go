package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rocky.my.id/git/h8-assignment-2/api/orders/requests"
	"rocky.my.id/git/h8-assignment-2/api/orders/responses"
	"rocky.my.id/git/h8-assignment-2/models"
)

// Update godoc
// @Summary     Update order
// @Description Updates an order with the given ID and payload on the database.
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       id    path     int                        true "Order ID"
// @Param       order body     order_requests.OrderUpdate true "Update Order Payload"
// @Success     200   {object} order_responses.Data
// @Failure     400   {object} order_responses.Error
// @Failure     404   {object} order_responses.Error
// @Failure     500   {object} order_responses.Error
// @Router      /orders/{id} [put]
func (controller OrderController) Update(context *gin.Context) {
	var (
		id            = context.Param("id")
		order         models.Order
		orderUpdate   models.Order
		updatePayload order_requests.OrderUpdate
		err           error
	)

	if err := context.ShouldBind(&updatePayload); err != nil {
		context.JSON(
			http.StatusBadRequest,
			order_responses.Error{
				Message: err.Error(),
				Status:  "error",
			},
		)
		return
	}

	updatePayload.BindToModel(&orderUpdate)

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

	err = controller.Actions.UpdateOrder(&order, &orderUpdate)
	if err != nil {
		context.AbortWithStatusJSON(
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
		order_responses.Data{
			Data:    order,
			Message: "updated",
		},
	)
}
