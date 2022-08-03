package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-app/src/services"
)

func CreateUser(ctx *gin.Context) {
	var requestBody struct {
		Name string `json:"name"`
	}

	ctx.ShouldBindJSON(&requestBody)

	user := services.CreateUser(requestBody.Name)

	fmt.Println(requestBody.Name)

	ctx.JSON(200, user)
}