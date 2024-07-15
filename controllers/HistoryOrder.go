package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/services"
)

func GetOrderHistory(ctx *gin.Context) {
	var Client models.Client
	if err := ctx.ShouldBindJSON(&Client); err != nil {
		ctx.JSON(500, gin.H{
			"msg": "can't bind Client error: " + err.Error(),
		})
		return
	}
	result, err := services.GetOrderHistory(&Client)
	if err != nil {
		ctx.JSON(500, gin.H{
			"msg": "can't found" + err.Error(),
		})
		return
	}

	ctx.JSON(200, result)
}

func SaveOrderHistory(ctx *gin.Context) {
	var HO models.HistoryOrder
	var Client models.Client

	if err := ctx.ShouldBindJSON(&HO); err != nil {
		ctx.JSON(500, gin.H{
			"msg": "can't bind HistoryOrder" + err.Error(),
		})
		return

	}
	Client.Client_name = HO.Client_name
	Client.Exchange_name = HO.Exchange_name
	Client.Label = HO.Label
	Client.Pair = HO.Pair

	services.SaveClient(&Client)
	HO.ClientID = Client.ID
	services.SaveOrderHistory(&HO)

	ctx.JSON(200, gin.H{
		"msg": "saved",
		"id":  HO.ID,
	})
}
