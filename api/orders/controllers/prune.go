package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rockyprabowo/assignment-2/api/orders/responses"
)

func (controller OrderController) Prune(context *gin.Context) {
	count, err := controller.Actions.Prune()
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			order_responses.Error{
				Message: "Prune failed.",
				Status:  "error",
			},
		)
		return
	}
	context.JSON(
		http.StatusOK,
		order_responses.RowsAffected{
			Count:   int(count),
			Message: "Prune complete.",
		},
	)
}
