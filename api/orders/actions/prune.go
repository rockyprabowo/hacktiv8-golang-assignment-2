package order_actions

import . "rocky.my.id/git/h8-assignment-2/models"

// Prune prunes the database.
func (actions OrderActions) Prune() (int64, error) {
	var (
		affectedRows int64
		models       = []any{
			&Item{}, &Order{},
		}
	)
	for _, model := range models {
		pruneQuery := actions.Database.Where("1 = 1").Delete(model)

		if err := pruneQuery.Error; err != nil {
			return 0, err
		}
		affectedRows += pruneQuery.RowsAffected
	}
	actions.Database.Exec("ALTER SEQUENCE items_item_id_seq RESTART")
	actions.Database.Exec("ALTER SEQUENCE orders_order_id_seq RESTART")
	actions.Database.Exec("DROP TABLE orders CASCADE")

	if err := actions.Database.AutoMigrate(models...); err != nil {
		return affectedRows, err
	}
	return affectedRows, nil
}
