package models

import "gorm.io/gorm"

// OrderBook represents the full order book for a given exchange and pair
type OrderBook struct {
	gorm.Model   `gorm:"embedded"`
	ID           uint         `gorm:"column:id;autoincrement;primaryKey"`
	ExchangeName string       `gorm:"Key"`
	Pair         string       `gorm:"key"`
	Asks         []DepthOrderAsks `gorm:"foreignKey:order_book_id;references:ID"`
	Bids         []DepthOrderBids `gorm:"foreignKey:order_book_id;references:ID"`
}
