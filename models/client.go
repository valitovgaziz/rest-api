package models

// Clien клиент
type Client struct {
	ID            uint   `gorm:"column:id;autoincrement;primaryKey;not null" json:"id"`
	Client_name   string `gorm:"column:client_name;not null;type:varchar(255);uniqueIndex" json:"clientName"`
	Exchange_name string `gorm:"column:exchange_nam;not nulltype:varchar(255)" json:"exchangeName"`
	Label         string `gorm:"type:varchar(255)" json:"label"`
	Pair          string `gorm:"type:varchar(255)" json:"pair"`
}
