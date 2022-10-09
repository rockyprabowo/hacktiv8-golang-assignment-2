package order_requests

import (
	"rocky.my.id/git/h8-assignment-2/helpers/slices"
	. "rocky.my.id/git/h8-assignment-2/models"
)

// OrderUpdate
// @Description Represents a request payload for order update.
type OrderUpdate struct {
	CustomerName string       `form:"customerName" json:"customerName" example:"Marvin"` // This is the customer name
	Items        []ItemUpdate `form:"items" json:"items" binding:"required,dive"`        // This is the order items
}

// BindToModel binds the request to the actual data model.
func (from *OrderUpdate) BindToModel(to *Order) {
	to.CustomerName = from.CustomerName
	to.Items = slices.Map(from.Items, ItemUpdateMapper)
}
