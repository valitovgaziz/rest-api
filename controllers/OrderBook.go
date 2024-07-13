package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/services"
	"github.com/valyala/fastjson"
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
	var p fastjson.Parser
	Bids = p.Parse(ctx.Param("bids"))
	Asks = p.Parse(ctx.Param("asks"))
	bidss := []models.DepthOrderBids{}
	asdss := []models.DepthOrderAsks{}
	err := services.SaveOrderBook(ctx.Param("exchange_name"), ctx.Param("pair"), bidss, asdss)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "OrderBook saved",
	})
}
