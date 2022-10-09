package models

import "time"

// Order
// @Description Represent an order of a customer with the order date and one or more item(s).
type Order struct {
	ID           uint      `json:"orderId" gorm:"primaryKey;column:order_id" example:"1"`
	CustomerName string    `form:"customerName" json:"customerName" example:"Marvin"`
	OrderedAt    time.Time `form:"orderedAt" json:"orderedAt" gorm:"autoCreateTime" example:"2022-02-22T20:00:00.0000000+00:00"`
	Items        []Item    `form:"items" json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
