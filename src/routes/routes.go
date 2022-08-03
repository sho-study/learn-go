package routes

import (
	"github.com/gin-gonic/gin"
	"my-app/src/controllers"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

  router.POST("/user", controllers.CreateUser)

	return router
}
