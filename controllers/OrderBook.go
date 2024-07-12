package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetOrderBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetOrderBook",
	})
}

func SaveOrderBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "SaveOrderBook",
	})
}
