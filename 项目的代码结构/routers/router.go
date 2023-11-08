package routers

import (
	"github.com/gin-gonic/gin"
	"hello/项目的代码结构/Controller"
	"hello/项目的代码结构/Model"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", Model.CreateTodo)
		v1Group.GET("/todo/:id", Controller.GetTodo)
		v1Group.PUT("/todo/:id", Controller.PutTodo)
		v1Group.POST("/todo", Controller.PostTodo)
		v1Group.DELETE("/todo", Controller.DeleteTodo)
	}
	return r
}
