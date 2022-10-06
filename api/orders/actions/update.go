package order_actions

import (
	"gorm.io/gorm"
	"rockyprabowo/assignment-2/helpers/slices"
	. "rockyprabowo/assignment-2/models"
)

func (actions OrderActions) UpdateOrder(order, updatePayload *Order) error {
	originalOrderItemIDs := slices.Map(order.Items, ItemMapID)
	order.Items = updatePayload.Items

	err := actions.Database.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(order).
		Error
	if err != nil {
		return err
	}

	savedOrderItemIDs := slices.Map(order.Items, ItemMapID)

	deletedItemIDs := slices.Diff(originalOrderItemIDs, savedOrderItemIDs)

	actions.Database.Delete(&Item{}, deletedItemIDs)
	return nil
}
