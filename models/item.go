package models

// Item represents an Item tied to an order.
type Item struct {
	ID          uint   `json:"lineItemId" gorm:"primaryKey;column:item_id"`
	ItemCode    string `form:"itemCode" json:"itemCode"`
	Description string `form:"description" json:"description"`
	Quantity    uint   `form:"quantity" json:"quantity"`
	OrderID     uint   `json:"-"`
	Order       Order  `json:"-"`
}

// ItemMapID is a function for a mapper that returns an Item ID.
func ItemMapID(item Item) uint {
	return item.ID
}

// ItemFilterIDZero is a function that filters item with ID = 0
// Unused for now.
// noinspection GoUnusedExportedFunction
func ItemFilterIDZero(item Item) bool {
	return item.ID == 0
}

// ItemFilterIDNonZero is a function that filters item with non-zero ID
// Unused for now.
// noinspection GoUnusedExportedFunction
func ItemFilterIDNonZero(item Item) bool {
	return item.ID != 0
}
