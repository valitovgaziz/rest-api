package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/initializers"
	"github.com/valitovgaziz/rest-api/migrate"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load environment variables", err)
	}
	sqlDB, err := initializers.DB.DB()
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}
	defer sqlDB.Close()

	migrate.Migrate()

	router := server.Group("/api")
	router.GET("/hello", func(c *gin.Context) {
		message := "from gin"
		c.JSON(http.StatusOK, gin.H{"status": "succes", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
