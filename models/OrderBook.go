package models

// OrderBook represents the full order book for a given exchange and pair
type OrderBook struct {
    ExchangeName string			`gorm:"keyvalue"`
    Pair         string
    Asks         []DepthOrder
    Bids         []DepthOrder
}