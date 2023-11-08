package Model

import (
	"github.com/gin-gonic/gin"
	"hello/项目的代码结构/DB"
)

// Todo Model
type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// todo 的增删改查
func CRUD() {
	var db Todo
	DB.DB.First(&db)
}

func CreateTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	//...
}
