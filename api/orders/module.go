package order_module

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rockyprabowo/assignment-2/api/orders/actions"
	"rockyprabowo/assignment-2/api/orders/controllers"
	"rockyprabowo/assignment-2/api/orders/routes"
	"rockyprabowo/assignment-2/routers"
)

type OrderModule struct {
	Routes     routers.RouterContract
	Controller order_controllers.OrderControllerContract
	Actions    order_actions.OrderActionsContract
}

func NewOrderModule(
	actions order_actions.OrderActionsContract,
	controller order_controllers.OrderControllerContract,
	routes routers.RouterContract,
) *OrderModule {
	return &OrderModule{
		Actions:    actions,
		Controller: controller,
		Routes:     routes,
	}
}

func SetupDefault(engine *gin.Engine, db *gorm.DB) {
	orderActions := order_actions.NewOrderActions(db)
	orderController := order_controllers.NewOrdersController(db, orderActions)
	orderRoutes := order_routes.NewOrderRoutes(engine, orderController)

	module := NewOrderModule(
		orderActions,
		orderController,
		orderRoutes,
	)

	module.Routes.Setup()
}
