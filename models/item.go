package models

type Item struct {
	ID          uint   `json:"lineItemId" gorm:"primaryKey;column:item_id"`
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
	OrderID     uint   `json:"-"`
	Order       Order  `json:"-"`
}

func ItemMapID(item Item) uint {
	return item.ID
}

// noinspection GoUnusedExportedFunction
func ItemFilterIDZero(item Item) bool {
	return item.ID == 0
}

// noinspection GoUnusedExportedFunction
func ItemFilterIDNonZero(item Item) bool {
	return item.ID != 0
}
