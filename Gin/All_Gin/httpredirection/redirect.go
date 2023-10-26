package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//永久重定向
	r := gin.Default()
	r.GET("/demo", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"msg":"ok",
		//})
		c.Redirect(http.StatusMovedPermanently, "https://git-scm.com/")
	})

	//临时重定向
	r.GET("/a", func(c *gin.Context) {
		//跳转到/b 对应的路由处理函数
		c.Request.URL.Path = "/b" //将请求的URL 修改
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	r.Run(":8080")
}
