package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rocky.my.id/git/h8-assignment-2/api/orders/requests"
	"rocky.my.id/git/h8-assignment-2/http/responses"
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
// @Success     200   {object} responses.WithSingleData[models.Order]
// @Failure     400   {object} responses.Error
// @Failure     404   {object} responses.Error
// @Failure     500   {object} responses.Error
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
			responses.Error{
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
			responses.Error{
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
			responses.Error{
				Message: err.Error(),
				Status:  "error",
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		responses.WithSingleData[models.Order]{
			Data:   order,
			Status: "updated",
		},
	)
}
