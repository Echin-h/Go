package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"` // `json:"username"`
	Address  string `form:"address"`  //`json:"address"`
}

func main() {
	r := gin.Default()
	r.GET("/todo", func(c *gin.Context) {
		var u UserInfo
		//err := c.ShouldBindJSON(&u)
		//if err != nil {
		//	log.Panicln(err)
		//}
		err := c.ShouldBind(&u)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(u)
		c.JSON(http.StatusOK, u)
	})
	r.Run(":8080")
}

//shouldbindjson 一般来说只能处理json数据，建议处理json数据的时候直接使用它
//shouldbind 一般来说可以处理多个类型，postform,querystring等

//注意：   1.传递指针  2.结构体首字母大写并且加标签
//因为要改变u的值，所以需要传入指针
/* 用ShouldBind函数的时候我不知道用户会传进来一个什么值，所以
 而 u 是结构体，不知道里面有啥字段，所以需要用标签来标记，
要让外面的人拿到里面的字段，所以里面的字段首字母需要大写
*/
