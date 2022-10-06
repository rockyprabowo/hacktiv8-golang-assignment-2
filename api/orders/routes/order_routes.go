package order_routes

import (
	"github.com/gin-gonic/gin"
	"rockyprabowo/assignment-2/api/orders/controllers"
)

type OrderRoutes struct {
	Engine  *gin.Engine
	Handler order_controllers.OrderControllerContract
}

func NewOrderRoutes(
	engine *gin.Engine,
	handler order_controllers.OrderControllerContract,
) *OrderRoutes {
	return &OrderRoutes{Engine: engine, Handler: handler}
}

func (r OrderRoutes) Setup() {
	orderRoutes := r.Engine.Group("/orders")
	{
		orderRoutes.GET("", r.Handler.GetAll)
		orderRoutes.POST("", r.Handler.Create)
		orderRoutes.GET("/:id", r.Handler.GetById)
		orderRoutes.PUT("/:id", r.Handler.Update)
		orderRoutes.DELETE("/:id", r.Handler.Delete)
		orderRoutes.POST("/prune", r.Handler.Prune)
	}
}
