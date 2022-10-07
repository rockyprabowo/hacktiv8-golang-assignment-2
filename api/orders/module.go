package order_module

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rockyprabowo/assignment-2/api/orders/actions"
	"rockyprabowo/assignment-2/api/orders/controllers"
	"rockyprabowo/assignment-2/api/orders/routes"
	"rockyprabowo/assignment-2/routers"
)

// OrderModule defines the order module.
type OrderModule struct {
	Routes     routers.RouterContract
	Controller order_controllers.OrderControllerContract
	Actions    order_actions.OrderActionsContract
}

// SetupDefault instantiates a "default" instance of OrderModule and set the routes up.
// TODO: Implement/use a package for Dependency Injection Container?
func SetupDefault(engine *gin.Engine, db *gorm.DB) {
	orderActions := order_actions.NewOrderActions(db)
	orderController := order_controllers.NewOrdersController(db, orderActions)
	orderRoutes := order_routes.NewOrderRoutes(engine, orderController)

	module := OrderModule{
		Routes:     orderRoutes,
		Controller: orderController,
		Actions:    orderActions,
	}

	module.Routes.Setup()
}
