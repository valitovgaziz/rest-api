package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/storage"
)

var (
	server *gin.Engine
)

func init() {
	config, err := storage.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load environment variables", err)
	}

	storage.ConnectDB(&config)

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
