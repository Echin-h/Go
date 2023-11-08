package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 存放具体实现的代码

func PostTodo(c *gin.Context) {
	panic("haha")
}
func DeleteTodo(c *gin.Context) {
	panic("haha")
}
func PutTodo(c *gin.Context) {
	panic("haha")
}
func GetTodo(c *gin.Context) {
	panic("haha")
}
