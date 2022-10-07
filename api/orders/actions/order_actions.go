package order_actions

import (
	"gorm.io/gorm"
	. "rockyprabowo/assignment-2/models"
)

// OrderActionsContract defines a contract for order related actions.
type OrderActionsContract interface {
	CreateOrder(newOrder *Order) error
	GetAllOrders() ([]Order, error)
	GetOrderById(id string) (Order, error)
	GetOrdersWith(associations string, args ...any) ([]Order, error)
	OrderExists(id string) (bool, error)
	UpdateOrder(order, updateOrder *Order) error
	DeleteOrder(order *Order) (int64, error)
	DeleteOrderById(id string) (int64, error)
	Prune() (int64, error)
}

// OrderActions defines a concrete implementations of order actions.
type OrderActions struct {
	Database *gorm.DB
}

// NewOrderActions instantiate an OrderActions.
func NewOrderActions(db *gorm.DB) OrderActionsContract {
	return &OrderActions{Database: db}
}
