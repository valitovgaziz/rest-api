package models

// Clien клиент
type Client struct {
	ID            uint   `gorm:"column:id;autoincrement;primaryKey"`
	Client_name   string `gorm:"column:client_name;not null;type:varchar(255);uniqueIndex"`
	Exchange_name string `gorm:"column:exchange_nam;not nulltype:varchar(255)"`
	Label         string `gorm:"type:varchar(255)"`
	Pair          string `gorm:"type:varchar(255)"`
}
