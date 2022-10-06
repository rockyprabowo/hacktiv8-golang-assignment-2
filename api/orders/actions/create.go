package order_actions

import . "rockyprabowo/assignment-2/models"

func (actions OrderActions) CreateOrder(newOrder *Order) error {
	err := actions.Database.Create(newOrder).Error
	return err
}
