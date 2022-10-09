package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rocky.my.id/git/h8-assignment-2/http/responses"
	"rocky.my.id/git/h8-assignment-2/models"
)

// GetAll godoc
// @Summary     Fetch all orders
// @Description Fetch all order(s) from the database.
// @Tags        orders
// @Produce     json
// @Success     200 {object} responses.WithDataList[models.Order]
// @Failure     500 {object} responses.Error
// @Router      /orders [get]
func (controller OrderController) GetAll(context *gin.Context) {
	var (
		orders []models.Order
		err    error
	)

	if orders, err = controller.Actions.GetAllOrders(); err != nil {
		context.JSON(
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
		responses.WithDataList[models.Order]{
			Data:    orders,
			Count:   len(orders),
			Message: "success",
		},
	)
}
