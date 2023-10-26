package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string
	Chouhao string
	Age     string // query传参是字符串类型的
}

func main() { //任何请求开头的 都是用c.开头的
	r := gin.Default()
	//查询的时候遵循  样例： ?query=""&age=18  的原则=>  也就是 key = value 的原则
	r.GET("/demo", func(c *gin.Context) {
		//获取浏览器那边请求携带的 query string 参数
		query := c.Query("query") //   左边query表示变量，是样例中右边引号中的内容，而右边Query中的query是左边样例中的query
		age := c.Query("age")
		//query := c.DefaultQuery("query", "somebody") //这个就表示，如果query不存在则显示somebody
		/*query, ok := c.GetQuery("query")
		if !ok {
			query = "halo"
		}*/
		person := Person{
			Name:    query,
			Chouhao: "笨蛋",
			Age:     age,
		}
		c.JSON(200, person)
	})

	r.Run(":8080")
}
