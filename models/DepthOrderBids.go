package models

// DepthOrder represents a single order for bids in the order book
type DepthOrderBids struct {
	ID          uint    `gorm:"column:id;autoincrement;primaryKey" json:"id"`
	OrderBookId uint    `gorm:"column:order_book_id;not null" json:"order_book_id"`
	Price       float64 `gorm:"type:float4;not null" json:"price"`
	BaseQty     float64 `gorm:"type:float4;not null" json:"base_qty"`
}
