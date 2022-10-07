package order_routes

import (
	"github.com/gin-gonic/gin"
	"rocky.my.id/git/h8-assignment-2/api/orders/controllers"
)

// OrderRoutes defines a concrete implementation of RouterContract related to order.
type OrderRoutes struct {
	Engine  *gin.Engine
	Handler order_controllers.OrderControllerContract
}

// NewOrderRoutes instantiates an OrderRoutes.
func NewOrderRoutes(
	engine *gin.Engine,
	handler order_controllers.OrderControllerContract,
) *OrderRoutes {
	return &OrderRoutes{Engine: engine, Handler: handler}
}

// Setup defines the routes and its handler related to order.
func (r OrderRoutes) Setup() {
	orderRoutes := r.Engine.Group("/orders")
	{
		orderRoutes.GET("", r.Handler.GetAll)
		orderRoutes.POST("", r.Handler.Create)
		orderRoutes.GET("/:id", r.Handler.GetById)
		orderRoutes.PUT("/:id", r.Handler.Update)
		orderRoutes.DELETE("/:id", r.Handler.Delete)
		orderRoutes.POST("/__prune__", r.Handler.Prune)
	}
}
