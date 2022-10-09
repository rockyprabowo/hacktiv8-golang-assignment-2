package order_requests

import (
	. "rocky.my.id/git/h8-assignment-2/models"
)

// ItemUpdate
// @Description Represents a request payload for an item tied to an order create request.
type ItemUpdate struct {
	ID          uint   `json:"lineItemId" example:"1"`                                                          // This is the item ID
	ItemCode    string `form:"itemCode" json:"itemCode" binding:"required" example:"STUFF-1"`                   // This is the item code
	Description string `form:"description" json:"description" binding:"required" example:"A (new) description"` // This is the item description
	Quantity    uint   `form:"quantity" json:"quantity" binding:"required" example:"1337"`                      // This is the item quantity
}

// ItemUpdateMapper is a mapper function to the actual data model.
func ItemUpdateMapper(item ItemUpdate) (boundItem Item) {
	item.BindToModel(&boundItem)
	return
}

// BindToModel binds the request to the actual data model.
func (from *ItemUpdate) BindToModel(to *Item) {
	to.ID = from.ID
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = from.Quantity
}
