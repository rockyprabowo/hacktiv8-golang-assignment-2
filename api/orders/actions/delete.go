package order_actions

import . "rockyprabowo/assignment-2/models"

// DeleteOrder deletes an order with the given order.
func (actions OrderActions) DeleteOrder(order *Order) (int64, error) {
	result := actions.Database.Delete(order)
	return result.RowsAffected, result.Error
}

// DeleteOrderById deletes an order with the given ID.
func (actions OrderActions) DeleteOrderById(id string) (int64, error) {
	result := actions.Database.Delete(&Order{}, id)
	return result.RowsAffected, result.Error
}
