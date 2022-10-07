package order_requests

import (
	"rockyprabowo/assignment-2/helpers/slices"
	. "rockyprabowo/assignment-2/models"
)

// OrderCreate represents a request payload for order creation.
type OrderCreate struct {
	CustomerName string         `form:"customerName" json:"customerName" binding:"required"`
	Items        []ItemOnCreate `form:"items" json:"items" binding:"required"`
}

// BindToModel binds the request to the actual data model.
func (from OrderCreate) BindToModel(to *Order) {
	to.CustomerName = from.CustomerName
	to.Items = slices.Map(from.Items, ItemOnCreateMapper)
}