package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/storage"
)

var server *gin.Engine

func init() {
	// set default logger

	// load config from environment variables
	config, err := storage.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load environment variables", err)
	} // if error message

	// connect to PostgreSQL
	storage.ConnectDB(&config)

	// set server instance
	server = gin.Default()
}

func main() {
	config, err := storage.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load environment variables", err)
	}
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}
	defer storage.DropTables()
	defer storage.Close()

	router := server.Group("/api")
	router.GET("/hello", func(c *gin.Context) {
		message := "from gin"
		c.JSON(http.StatusOK, gin.H{"status": "succes", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
