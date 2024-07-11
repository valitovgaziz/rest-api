package models

import "gorm.io/gorm"

// DepthOrder represents a single order for bids in the order book
type DepthOrderBids struct {
	gorm.Model  `gorm:"embedded"`
	ID          uint    `gorm:"column:id;autoincrement;primaryKey"`
	OrderBookID uint    `gorm:"column:order_book_id;not null"`
	Price       float64 `gorm:"type:varchar(255);not null"`
	BaseQty     float64 `gorm:"type:varchar(255);not null"`
}