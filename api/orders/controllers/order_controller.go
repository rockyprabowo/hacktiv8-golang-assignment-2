package order_controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rockyprabowo/assignment-2/api/orders/actions"
)

// OrderControllerContract defines a contract for handling requests for order
type OrderControllerContract interface {
	GetById(*gin.Context)
	GetAll(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Prune(*gin.Context)
}

// OrderController defines a concrete implementations of order request handlers
type OrderController struct {
	Actions  order_actions.OrderActionsContract
	Database *gorm.DB
}

// NewOrdersController instantiate an OrderController
func NewOrdersController(
	db *gorm.DB,
	actions order_actions.OrderActionsContract,
) OrderControllerContract {
	return &OrderController{
		Database: db,
		Actions:  actions,
	}
}
