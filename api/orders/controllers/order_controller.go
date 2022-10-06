package order_controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rockyprabowo/assignment-2/api/orders/actions"
)

type OrderControllerContract interface {
	GetById(*gin.Context)
	GetAll(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Prune(*gin.Context)
}

type OrderController struct {
	Actions  order_actions.OrderActionsContract
	Database *gorm.DB
}

func NewOrdersController(db *gorm.DB, actions order_actions.OrderActionsContract) *OrderController {
	return &OrderController{Database: db, Actions: actions}
}
