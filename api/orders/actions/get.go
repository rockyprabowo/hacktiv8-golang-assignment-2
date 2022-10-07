package order_actions

import (
	"gorm.io/gorm/clause"
	. "rockyprabowo/assignment-2/models"
)

// OrderExists checks the existence of an order by its ID.
func (actions OrderActions) OrderExists(id string) (exist bool, err error) {
	result := actions.Database.Select("order_id").First(&Order{}, id)
	return result.RowsAffected > 0, result.Error
}

// GetOrderById fetches an order from the database by its ID.
func (actions OrderActions) GetOrderById(id string) (Order, error) {
	var order Order
	err := actions.Database.Preload(clause.Associations).First(&order, id).Error
	return order, err
}
