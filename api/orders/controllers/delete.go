package order_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller OrderController) Delete(context *gin.Context) {
	var (
		id      = context.Param("id")
		exist   bool
		deleted int64
		err     error
	)

	exist, err = controller.Actions.OrderExists(id)
	if !exist || err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	deleted, err = controller.Actions.DeleteOrderById(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"count": deleted, "message": "Deleted"})
}
