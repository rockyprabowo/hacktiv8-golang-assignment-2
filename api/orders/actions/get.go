package order_actions

import (
	"gorm.io/gorm/clause"
	. "rockyprabowo/assignment-2/models"
)

func (actions OrderActions) OrderExists(id string) (exist bool, err error) {
	result := actions.Database.Select("order_id").First(&Order{}, id)
	return result.RowsAffected > 0, result.Error
}

func (actions OrderActions) GetOrderById(id string) (Order, error) {
	var order Order
	err := actions.Database.Preload(clause.Associations).First(&order, id).Error
	return order, err
}
