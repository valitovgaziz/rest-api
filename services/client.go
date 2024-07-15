package services

import (
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

func SaveClient(client *models.Client) {
	storage.DB.Save(client)
}

func GetClient(id uint) *models.Client {
	var client models.Client
	storage.DB.First(&client, id)
	return &client
}
