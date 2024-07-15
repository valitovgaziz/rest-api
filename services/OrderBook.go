package services

import (
	"database/sql"

	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

func GetOrderBook(exchange_name string, pair string) ([]*models.OrderBook, error) {
	var orderBooks []*models.OrderBook

	// Используем метод Find для поиска записей в базе данных
	if err := storage.DB.Where("exchange_name = ? AND pair = ?", exchange_name, pair).Find(&orderBooks).Error; err != nil {
		if err == sql.ErrNoRows {
			return []*models.OrderBook{}, nil // Возвращаем пустой массив и nil, чтобы указать на отсутствие записей
		}
		return nil, err // Возвращаем ошибку, если произошла другая ошибка
	}

	return orderBooks, nil
}

func SaveOrderBook(OB models.OrderBook) error {
	/*
	for _, bid := range OB.Bids {
		if err := storage.DB.Create(bid).Error; err != nil {
			return err
		}
	}
	for _, ask := range OB.Asks {
		if err := storage.DB.Create(ask).Error; err != nil {
			return err
		}
	}
		*/
	

	if err := storage.DB.Create(&OB).Error; err != nil {
		return err
	}
	return nil
}