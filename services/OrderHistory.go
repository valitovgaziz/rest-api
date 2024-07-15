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
	if err := storage.DB.Where("client_id = ?", client.ID).Find(&historyOrders).Error; err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no orders found for client by client_id")
		}
		return nil, err
	}

	return historyOrders, nil
}

func SaveOrderHistory(order *models.HistoryOrder) error {

	if err := storage.DB.Save(&order).Error; err != nil {
		return err
	}
	return nil
}
