package models

// OrderBook represents the full order book for a given exchange and pair
type OrderBook struct {
	ID           int64            `gorm:"primary_key;autoincrement;column:id"`
	ExchangeName string           `gorm:"column:exchange_name;not null" json:"exchange_name"`
	Pair         string           `gorm:"column:pair;not null" json:"pair"`
	Asks         []DepthOrderAsks `gorm:"foreignKey:order_book_id;references:ID" json:"asks"`
	Bids         []DepthOrderBids `gorm:"foreignKey:order_book_id;references:ID" json:"bids"`
}
