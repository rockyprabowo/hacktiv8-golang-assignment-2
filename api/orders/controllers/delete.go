package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rocky.my.id/git/h8-assignment-2/api/orders/responses"
)

// Delete godoc
// @Summary     Delete order
// @Description Deletes an order with the given ID from the database.
// @Tags        orders
// @Produce     json
// @Param       id  path     int true "Order ID"
// @Success     200 {object} order_responses.RowsAffected
// @Failure     404 {object} order_responses.Error
// @Failure     500 {object} order_responses.Error
// @Router      /orders/{id} [delete]
func (controller OrderController) Delete(context *gin.Context) {
	var (
		id      = context.Param("id")
		exist   bool
		deleted int64
		err     error
	)

	exist, err = controller.Actions.OrderExists(id)
	if !exist || err != nil {
		context.JSON(
			http.StatusNotFound,
			order_responses.Error{
				Message: err.Error(),
				Status:  "error",
			},
		)
		return
	}

	deleted, err = controller.Actions.DeleteOrderById(id)
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
		order_responses.RowsAffected{
			Count:   int(deleted),
			Message: "deleted",
		},
	)
}
