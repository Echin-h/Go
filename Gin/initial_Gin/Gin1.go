package main

import "github.com/gin-gonic/gin"

func sayhello(c *gin.Context) { //gin.context用于管理项目全局的设计和配置
	c.JSON(200, gin.H{
		"message": "hello",
		"name":    "ren",
	}) //JSON响应
}

func main() {
	r := gin.Default()
	// URL后面写函数，匿名函数也行
	r.GET("/hello", sayhello)
	//启动服务

	r.Run(":8080")
}
