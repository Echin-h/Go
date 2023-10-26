package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()

	r.GET("/demo", func(c *gin.Context) {

	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		file, err := c.FormFile("f1") //从请求中携带的参数一样
		if err != nil {
			fmt.Println(err)
			return
		} else {
			dst := path.Join("./", file.Filename) //路径选择
			c.SaveUploadedFile(file, dst)         //前面的是文件， 后面的是路径
			//form, err := c.MultipartForm()  多个文件上传
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}
		//将读取到的文件保存到本地（服务端）

	})

	r.Run(":8080")
}
