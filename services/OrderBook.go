package services

import (
	"database/sql"

	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

func GetOrderBook(exchange_name string, pair string) ([]*models.OrderBook, error) {
	var orderBooks []*models.OrderBook

	// Используем метод Find для поиска записей в базе данных
	if err := storage.DB.Where("exchange = ? AND pair = ?", exchange_name, pair).Find(&orderBooks).Error; err != nil {
		if err == sql.ErrNoRows {
			return []*models.OrderBook{}, nil // Возвращаем пустой массив и nil, чтобы указать на отсутствие записей
		}
		return nil, err // Возвращаем ошибку, если произошла другая ошибка
	}

	return orderBooks, nil
}

func SaveOrderBook(OB models.OrderBook) error {
	if err := storage.DB.Save(OB).Error; err != nil {
		return err
	}
	return nil
}
