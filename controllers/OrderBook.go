package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/services"
)

func GetOrderBook(ctx *gin.Context) {
	OB, err := services.GetOrderBook(ctx.Param("exchange_name"), ctx.Param("pair"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, OB)
}

func SaveOrderBook(ctx *gin.Context) {
	var OB models.OrderBook
	err := ctx.ShouldBindJSON(&OB)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err1 := services.SaveOrderBook(OB)
	if err1 != nil {
		ctx.JSON(400, gin.H{"error": err1.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "OrderBook saved",
	})
}
