package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rocky.my.id/git/h8-assignment-2/http/responses"
)

func (controller OrderController) Prune(context *gin.Context) {
	count, err := controller.Actions.Prune()
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			responses.Error{
				Message: "Prune failed.",
				Status:  "error",
			},
		)
		return
	}
	context.JSON(
		http.StatusOK,
		responses.WithRowsAffected{
			Count:   int(count),
			Message: "Prune complete.",
		},
	)
}
