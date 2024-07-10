package models

import "time"

// HistoryOrder represents a single trade history record
type HistoryOrder struct {
    ClientID    string
    TradeID     string
    Timestamp   time.Time
    Exchange    string
    Pair        string
    Side        string // BUY or SELL
    Price       float64
    Quantity    float64
    TotalAmount float64
    Fee         float64
    Status      string // FILLED, CANCELLED, etc.
}
