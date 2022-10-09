package order_actions

import (
	"gorm.io/gorm"
	"rocky.my.id/git/h8-assignment-2/helpers/options"
	"rocky.my.id/git/h8-assignment-2/helpers/slices"
	. "rocky.my.id/git/h8-assignment-2/models"
)

func vetOrderItems(original []Item, update []Item) (upsert []Item, deleted []uint) {
	originalOrderItemIDs := slices.Map(original, ItemMapID)
	updateOrderItemIDs := slices.Map(update, ItemMapID)

	updatableIDs := slices.Intersect(originalOrderItemIDs, updateOrderItemIDs)
	upsert = slices.Filter(update, func(i Item) bool {
		return slices.Contains(updatableIDs, i.ID) || i.ID == 0
	})
	deleted = slices.Diff(originalOrderItemIDs, updatableIDs)

	return
}

// UpdateOrder updates an order on the database.
func (actions OrderActions) UpdateOrder(order, updatePayload *Order) error {
	upsertItems, deletedItemIDs := vetOrderItems(order.Items, updatePayload.Items)
	order.CustomerName = options.Default(updatePayload.CustomerName, order.CustomerName)
	order.Items = upsertItems

	err := actions.
		Database.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(order).
		Error
	if err != nil {
		return err
	}

	if len(deletedItemIDs) > 0 {
		actions.Database.Delete(&Item{}, deletedItemIDs)
	}

	return nil
}
