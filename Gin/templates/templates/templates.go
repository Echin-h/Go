package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type msg struct {
	Name string `json:"hello"`
	Age  int
}

func main() {
	r := gin.Default()
	data := map[string]interface{}{
		"饭饭": "吃饭",
		"朱哥": "吃草",
		"年纪": 123456789,
	}

	m := msg{
		Name: "沙河",
		Age:  18,
	}
	r.GET("/demo", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{ //json序列化 是默认其他包,所以结构体要大写
			"map":    data,
			"struct": m,
		})

	})
	err := r.Run(":8080")
	if err != nil {
		log.Panicln("run is wrong", err)
		return
	}
}
