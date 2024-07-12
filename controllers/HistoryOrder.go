package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/services"
	"github.com/valitovgaziz/rest-api/storage"
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

	if err := storage.DB.Save(&OH); err != nil {
		ctx.JSON(500, gin.H{
			"msg": "can't save",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "saved",
		"id":  OH.ID,
	})
}
