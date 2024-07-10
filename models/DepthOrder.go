package models


// DepthOrder represents a single order in the order book
type DepthOrder struct {
    Price  float64	`gorm:"type:varchar(255);not null"`
    BaseQty float64	`gorm:"type:varchar(255);not null"`
}