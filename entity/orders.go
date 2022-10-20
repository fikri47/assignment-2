package entity

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt,omitempty"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}
