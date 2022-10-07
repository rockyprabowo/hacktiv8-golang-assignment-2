package order_requests

import (
	. "rockyprabowo/assignment-2/models"
)

// ItemOnUpdate represents a request payload for an item tied to an order create request.
type ItemOnUpdate struct {
	ID          uint   `json:"lineItemId" binding:"optional"`
	ItemCode    string `form:"itemCode" json:"itemCode" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Quantity    uint   `form:"quantity" json:"quantity" binding:"required"`
}

// ItemOnUpdateMapper is a mapper function to the actual data model.
func ItemOnUpdateMapper(item ItemOnUpdate) (boundItem Item) {
	item.BindToModel(&boundItem)
	return
}

// BindToModel binds the request to the actual data model.
func (from ItemOnUpdate) BindToModel(to *Item) {
	to.ID = from.ID
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = from.Quantity
}
