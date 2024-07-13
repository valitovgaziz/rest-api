package models

// DepthOrder represents a single order for asks in the order book
type DepthOrderAsks struct {
	ID          uint    `gorm:"column:id;autoincrement;primaryKey"`
	OrderBookId uint    `gorm:"column:order_book_id;not null"`
	Price       float64 `gorm:"type:varchar(255);not null"`
	BaseQty     float64 `gorm:"type:varchar(255);not null"`
}
