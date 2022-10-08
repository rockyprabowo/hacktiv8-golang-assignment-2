package models

// Item
// @Description Represents an item tied to an order.
type Item struct {
	ID          uint   `json:"lineItemId" gorm:"primaryKey;column:item_id"  example:"1"`
	ItemCode    string `form:"itemCode" json:"itemCode" example:"STUFF-1"`
	Description string `form:"description" json:"description" example:"A description"`
	Quantity    uint   `form:"quantity" json:"quantity" example:"1337"`
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
