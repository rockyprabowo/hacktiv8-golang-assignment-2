package order_actions

import (
	"gorm.io/gorm"
	. "rockyprabowo/assignment-2/models"
)

type OrderActionsContract interface {
	CreateOrder(newOrder *Order) error
	GetAllOrders() ([]Order, error)
	GetOrderById(id string) (Order, error)
	OrderExists(id string) (bool, error)
	UpdateOrder(order, updateOrder *Order) error
	DeleteOrder(order *Order) (int64, error)
	DeleteOrderById(id string) (int64, error)
	Prune() (int64, error)
}

type OrderActions struct {
	Database *gorm.DB
}

func NewOrderActions(db *gorm.DB) *OrderActions {
	return &OrderActions{Database: db}
}
