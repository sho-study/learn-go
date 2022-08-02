package routes

import (
	"github.com/gin-gonic/gin"
	"my-app/src/controllers"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.GetHello)
	router.GET("/hoge", controllers.GetHello)

	return router
}
