package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rocky.my.id/git/h8-assignment-2/api/orders/requests"
	"rocky.my.id/git/h8-assignment-2/http/responses"
	"rocky.my.id/git/h8-assignment-2/models"
)

// Create godoc
// @Summary     Create order
// @Description Creates a new order and persists them to the database.
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       order body     order_requests.OrderCreate true "Create Order Request"
// @Success     200   {object} responses.WithData[models.Order]
// @Failure     400   {object} responses.Error
// @Failure     500   {object} responses.Error
// @Router      /orders [post]
func (controller OrderController) Create(context *gin.Context) {
	var (
		createPayload order_requests.OrderCreate
		order         models.Order
		err           error
	)
	if err := context.ShouldBind(&createPayload); err != nil {
		context.JSON(
			http.StatusBadRequest,
			responses.Error{
				Message: err.Error(),
				Status:  "error",
			},
		)
		return
	}

	createPayload.BindToModel(&order)

	if err = controller.Actions.CreateOrder(&order); err != nil {
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
		responses.WithData[models.Order]{
			Data:   order,
			Status: "created",
		},
	)
}
