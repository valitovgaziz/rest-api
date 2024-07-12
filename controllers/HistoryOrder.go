package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetOrderHistory(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success HistoryOrder",
	})
}

func SaveOrderHistory(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success SaveOrder",
	})
}

