package models

// OrderBook represents the full order book for a given exchange and pair
type OrderBook struct {
	ID           int64            `gorm:"primary_key;auto_increment;column:id"`
	ExchangeName string           `gorm:"column:exchange_name;not null"`
	Pair         string           `gorm:"column:pair;not null"`
	Asks         []DepthOrderAsks `gorm:"foreignKey:order_book_id;references:ID"`
	Bids         []DepthOrderBids `gorm:"foreignKey:order_book_id;references:ID"`
}
