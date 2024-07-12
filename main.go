package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/controllers"
	"github.com/valitovgaziz/rest-api/storage"
	"github.com/valitovgaziz/rest-api/configs"
)

var server *gin.Engine

func init() {
	// set default logger

	// load config from environment variables
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load environment variables", err)
	} // if error message

	// connect to PostgreSQL
	storage.ConnectDB(&config)

	// set server instance
	server = gin.Default()
}

func main() {
	// Загрузка переменных окружения
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load environment variables", err)
	}
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}
	defer storage.DropTables()
	defer storage.Close()

	// Маршрутизация
	OrderBookRouter := server.Group("/OrderBook")
	{
		OrderBookRouter.GET("/GetOrderBook", controllers.GetOrderBook)
		OrderBookRouter.POST("/SaveOrderBook", controllers.SaveOrderBook)
	}
	OrderHistoryRouter := server.Group("/OrderHistory")
	{
		OrderHistoryRouter.GET("/GetOrderHistory", controllers.GetOrderHistory)
		OrderHistoryRouter.POST("/SaveOrderHistory", controllers.SaveOrderHistory)
	}

	log.Fatal(server.Run(":" + config.ServerPort))
}
