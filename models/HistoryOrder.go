package models

import (
	"time"
)

// HistoryOrder represents a single trade history record
type HistoryOrder struct {
	ID                    uint      `gorm:"columnt:id;primary_key;autoincrement;not null;"`
	ClientID              uint      `gorm:"column:client_id;foreignKey;not null;"`
	Client_name           string    `gorm:"column:client_name;not null;type:varchar(255);uniqueIndex"`
	Exchange_name         string    `gorm:"column:exchange_nam;not nulltype:varchar(255)"`
	Label                 string    `gorm:"type:varchar(255)"`
	Pair                  string    `gorm:"type:varchar(255)"`
	Side                  string    `gorm:"column:side;type:varchar(255);default:0" json:"Side"`
	Type                  string    `gorm:"column:type;type:varchar(255);default:0" json:"Type"`
	Base_qty              float64   `gorm:"column:base_qty;type:integer;default:0" json:"BaseQty"`
	Price                 float64   `gorm:"column:price;type:float8;default:0" json:"Price"`
	Algorithm_name_placed string    `gorm:"column:algorithm_name_placed;type:varchar(255);default:0" json:"AlgorithmNamePlaced"`
	Lowest_sell_prc       float64   `gorm:"column:lowest_sell_prc;type:float8;default:0" json:"LowestSellPrc"`
	Highest_buy_prc       float64   `gorm:"column:highest_buy_prc;type:float8;default:0" json:"HighestBuyPrc"`
	Commission_quote_qty  float64   `gorm:"column:commission_quote_qty;type:float8;default:0" json:"CommissionQuoteQty"`
	Time_placed           time.Time `gorm:"column:time_placed;type:timestamp;default:0" json:"TimePlaced"` // time.Time
}
