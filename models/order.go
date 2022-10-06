package models

import "time"

type Order struct {
	ID           uint      `json:"orderId" gorm:"primaryKey;column:order_id"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"autoCreateTime"`
	Items        []Item    `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
