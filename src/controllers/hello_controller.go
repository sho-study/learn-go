package controllers

import "github.com/gin-gonic/gin"

func GetHello(ctx *gin.Context) {
	ctx.JSONP(200, gin.H{
		"message": "hello",
	})
}