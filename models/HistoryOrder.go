package models

import (
	"time"

	"gorm.io/gorm"
)

// HistoryOrder represents a single trade history record
type HistoryOrder struct {
	gorm.Model	`embedded`
	Client					uint
	side 					string
	_type 					string	`gorm:"type"`
	base_qty 				float64
	price 					float64
	algorithm_name_placed 	string
	lowest_sell_prc 		float64
	highest_buy_prc 		float64
	commission_quote_qty 	float64
	time_placed 			time.Time
}
