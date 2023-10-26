package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() { //三种方法，各自选择
	r := gin.Default()
	r.POST("/demo", func(c *gin.Context) {
		//username := c.PostForm("name")
		//username := c.DefaultPostForm("name", "你输入的key有点小问题")
		username, ok := c.GetPostForm("name")
		if !ok {
			log.Println("ok, your getpostform is wrong", ok)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})

	})

	r.Run(":8080")

}
