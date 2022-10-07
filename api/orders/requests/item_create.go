package order_requests

import (
	. "rocky.my.id/git/h8-assignment-2/models"
)

// ItemOnCreate represents a request payload for an item tied to an order create request.
type ItemOnCreate struct {
	ItemCode    string `form:"itemCode" json:"itemCode" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Quantity    uint   `form:"quantity" json:"quantity" binding:"required"`
}

// ItemOnCreateMapper is a mapper function to the actual data model.
func ItemOnCreateMapper(item ItemOnCreate) (boundItem Item) {
	item.BindToModel(&boundItem)
	return
}

// BindToModel binds the request to the actual data model.
func (from ItemOnCreate) BindToModel(to *Item) {
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = from.Quantity
}
