package services

import (
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

func GetOrderBook(exchange_name string, pair string) ([]*models.OrderBook, error) {
	var orderBooks []*models.OrderBook
	result := storage.DB.Where("exchange_name = ? AND pair = ?", exchange_name, pair).Find(&orderBooks)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderBooks, nil

	// Используем метод Find для поиска записей в базе данных
	/*
		orderBooks = storage.DB.Where(&models.OrderBook{ExchangeName: exchange_name, Pair: pair}).Find(&orderBooks)

		for _, orderBook := range orderBooks {
			asks := storage.DB.Where("order_book_id = ?", orderBook.ID).Find(&DepthOrderAsks)
			bids := storage.DB.Where("order_book_id = ?", orderBook.ID).Find(&DepthOrderBids)
			orderBook.Asks = asks
			orderBook.Bids = bids
		}

		return orderBooks, nil */
}

func SaveOrderBook(OB models.OrderBook) error {
	if err := storage.DB.Create(&OB).Error; err != nil {
		return err
	}
	return nil
}
