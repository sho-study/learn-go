package controllers

import "github.com/gin-gonic/gin"

func GetHello(ctx *gin.Context) {
	ctx.JSONP(200, gin.H{
		"message": "hello",
	})
}

type HelloController struct { }

func NewHelloController() *HelloController {
  return &HelloController{}
}

func (controller *HelloController) fuga() {

}

func hoge() {
  con := NewHelloController()
  con.fuga()
}