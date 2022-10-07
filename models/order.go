package models

import "time"

// Order represent an order of a customer with the order date and one or more item(s).
type Order struct {
	ID           uint      `json:"orderId" gorm:"primaryKey;column:order_id"`
	CustomerName string    `form:"customerName" json:"customerName"`
	OrderedAt    time.Time `form:"orderedAt" json:"orderedAt" gorm:"autoCreateTime"`
	Items        []Item    `form:"items" json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
