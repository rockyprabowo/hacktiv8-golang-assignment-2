package order_requests

import (
	. "rocky.my.id/git/h8-assignment-2/models"
)

// ItemCreate
// @Description Represents a request payload for an item tied to an order create request.
type ItemCreate struct {
	ItemCode    string `form:"itemCode" json:"itemCode" binding:"required" example:"STUFF-1"`             // This is the item code
	Description string `form:"description" json:"description" binding:"required" example:"A description"` // This is the item description
	Quantity    uint   `form:"quantity" json:"quantity" binding:"required" example:"10"`                  // This is the item quantity
}

// ItemCreateMapper is a mapper function to the actual data model.
func ItemCreateMapper(item ItemCreate) (boundItem Item) {
	item.BindToModel(&boundItem)
	return
}

// BindToModel binds the request to the actual data model.
func (from *ItemCreate) BindToModel(to *Item) {
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = from.Quantity
}
