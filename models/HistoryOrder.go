package models

import (
	"time"
)

// HistoryOrder represents a single trade history record
type HistoryOrder struct {
	ID                    uint      `gorm:"columnt:id;primary_key;autoincrement;not null;"`
	Client_name           string    `gorm:"column:client_name;type:varchar(255);default:0"`
	Exchange_name         string    `gorm:"column:exchange_name;type:varchar(255);default:0"`
	Label                 string    `gorm:"column:label;type:varchar(255);default:0"`
	Pair                  string    `gorm:"column:pair;type:varchar(255);default:0"`
	Side                  string    `gorm:"column:side;type:varchar(255);default:0"`
	Type                  string    `gorm:"column:type;type:varchar(255);default:0"`
	Base_qty              float64   `gorm:"column:base_qty;type:integer;default:0"`
	Price                 float64   `gorm:"column:price;type:float8;default:0"`
	Algorithm_name_placed string    `gorm:"column:algorithm_name_placed;type:varchar(255);default:0"`
	Lowest_sell_prc       float64   `gorm:"column:lowest_sell_prc;type:float8;default:0"`
	Highest_buy_prc       float64   `gorm:"column:highest_buy_prc;type:float8;default:0"`
	Commission_quote_qty  float64   `gorm:"column:commission_quote_qty;type:float8;default:0"`
	Time_placed           time.Time `gorm:"column:time_placed;type:timestamp;default:0"` // time.Time
}
