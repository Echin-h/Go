package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 用param传参时 URI 需要写成 :name， 在APIfox中需要写成{name}(APIfox会自动把：转化成{})
	//如果APIfox中无法传递中文参数，就在请求参数中的设置中的URL自动编码选择遵循WHATWG 即可
	r.GET("/demo/:name/:age", func(c *gin.Context) {
		//获取路径参数
		name := c.Param("name")
		age := c.Param("age")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})

	r.Run(":8080")
}
