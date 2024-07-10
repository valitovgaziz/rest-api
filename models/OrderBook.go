package models

import "gorm.io/gorm"

// OrderBook represents the full order book for a given exchange and pair
type OrderBook struct {
	gorm.Model   `gorm:"embedded"`
	ID           uint         `gorm:"primarykey; not null;"`
	ExchangeName string       `gorm:"Key"`
	Pair         string       `gorm:"key"`
	Asks         []DepthOrder `gorm:"foreignKey:id"`
	Bids         []DepthOrder `gorm:"foreignKey:id"`
}
