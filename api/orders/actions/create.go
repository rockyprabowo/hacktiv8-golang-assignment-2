package order_actions

import . "rocky.my.id/git/h8-assignment-2/models"

// CreateOrder creates an order and write them into the database.
func (actions OrderActions) CreateOrder(newOrder *Order) error {
	err := actions.Database.Create(newOrder).Error
	return err
}
