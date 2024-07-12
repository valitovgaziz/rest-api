package services

import (
	"database/sql"
	"fmt"

	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

// Get OrderHistory по имени клиента
func GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	var historyOrders []*models.HistoryOrder

	// Используем метод Find для поиска записей в базе данных
	if err := storage.DB.Where("client_name = ?", client.Client_name).Find(&historyOrders).Error; err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no orders found for client")
		}
		return nil, err
	}

	return historyOrders, nil
}

func SaveOrderHistory(client *models.Client, order *models.HistoryOrder) error {
	// Создаем транзакцию для безопасного выполнения операций
	tx := storage.DB.Begin()

	// Добавляем новый заказ
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback() // В случае ошибки откатываем транзакцию
		return err
	}

	// Обновляем клиента, добавляя новый заказ в историю
	if err := tx.Model(&client).Association("HistoryOrders").Append(order); err != nil {
		tx.Rollback() // В случае ошибки откатываем транзакцию
		return err
	}

	tx.Commit() // В случае ошибки откатываем транзакцию
	return nil
}
