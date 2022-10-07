package order_actions

import (
	"gorm.io/gorm/clause"
	. "rockyprabowo/assignment-2/models"
)

// GetAllOrders fetches every order from the database.
func (actions OrderActions) GetAllOrders() (orders []Order, err error) {
	err = actions.Database.Preload(clause.Associations).Find(&orders).Error
	return
}

// GetOrdersWith fetches every order with associations from the database.
func (actions OrderActions) GetOrdersWith(
	associations string,
	args ...any,
) (orders []Order, err error) {
	err = actions.Database.Preload(associations, args).Find(&orders).Error
	return
}
