package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//r.Any()  //所有都囊括其中  ，然后用switch case 选择一下、
	todo := r.Group("/demo")
	{
		todo.GET("/hehe", func(c *gin.Context) {
		})
		todo.POST("/lala", func(c *gin.Context) {
		})
		todo.DELETE("wawa", func(c *gin.Context) {
		})
		todo.PUT("gege", func(c *gin.Context) {
		})
	}

	r.Run(":8080")
}

//路由组可以嵌套  随你怎么写反正
