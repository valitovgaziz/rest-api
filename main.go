package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/initializers"
)

var (
	server *gin.Engine
)

func init() {
	if config, err := initializers.LoadConfig("."); err!= nil {
		log.Fatal("Can't load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	router := server.Group("/api")
	router.GET("/hello", func(c *gin.Context) {
		message := "from gin"
		c.JSON(http.StatusOK, gin.H{"status": "succes", "message": message})
	})

	log.Fatal(server.)
}
