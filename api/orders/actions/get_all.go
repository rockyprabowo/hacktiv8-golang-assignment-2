package order_actions

import (
	"gorm.io/gorm/clause"
	. "rockyprabowo/assignment-2/models"
)

func (actions OrderActions) GetAllOrders() (orders []Order, err error) {
	err = actions.Database.Preload(clause.Associations).Find(&orders).Error
	return
}

func (actions OrderActions) GetOrdersWith(
	associations string,
	args ...any,
) (orders []Order, err error) {
	err = actions.Database.Preload(associations, args).Find(&orders).Error
	return
}
