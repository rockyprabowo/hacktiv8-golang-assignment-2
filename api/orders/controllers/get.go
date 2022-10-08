package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rocky.my.id/git/h8-assignment-2/http/responses"
	"rocky.my.id/git/h8-assignment-2/models"
)

// GetById
// @Summary     Fetch single order
// @Description Fetch an order by its ID from the database.
// @Tags        orders
// @Produce     json
// @Param       id  path     int true "Order ID"
// @Success     200 {object} responses.WithSingleData[models.Order]
// @Failure     404 {object} responses.Error
// @Failure     500 {object} responses.Error
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
			Status: "found",
		},
	)
}
