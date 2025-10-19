package model

import (
	"gorm.io/gorm"
)

// Order represents a transaction order in Paring
type Order struct {
	gorm.Model
	UserID     uint      `gorm:"not null"` // From JWT claim
	TotalPrice float64   `gorm:"not null"`
	Status     string    `gorm:"default:'PENDING'"`
	Items      []OrderItem
}

// OrderItem links to Product Service
type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID string  `gorm:"not null"` // ID from Product Service (MongoDB ObjectId as string)
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"` // Snapshot price at order time
}