package models

type Client struct {
	id            uint `gorm:"primary_key; not null; autoincrement"`
	client_name   string
	exchange_name string
	label         string
	pair          string
}
