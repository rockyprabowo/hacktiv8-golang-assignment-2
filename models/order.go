package models

import "time"

type Order struct {
	ID           uint      `json:"orderId" gorm:"primaryKey;column:order_id"`
	CustomerName string    `form:"customerName" json:"customerName" binding:"required"`
	OrderedAt    time.Time `form:"orderedAt" json:"orderedAt" gorm:"autoCreateTime"`
	Items        []Item    `form:"items" json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" binding:"required"`
}
