package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/services"
)

func GetOrderHistory(ctx *gin.Context) {
	result, err := services.GetOrderBook(ctx.Param("exchange_name"), ctx.Param("pair"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"msg": "can't found",
		})
		return
	}

	ctx.JSON(200, result)
}

func SaveOrderHistory(ctx *gin.Context) {
	var OH models.HistoryOrder
	var Client models.Client
	err1 := ctx.ShouldBindJSON(&Client)
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"msg": "can't bind client",
		})
		return
	}
	err2 := ctx.ShouldBindJSON(OH)
	if err2 != nil {
		ctx.JSON(500, gin.H{
			"msg": "can't bind HistoryOrder",
		})
		return

	}

	services.SaveOrderHistory(&Client, &OH)

	ctx.JSON(200, gin.H{
		"msg": "saved",
		"id":  OH.ID,
	})
}
