package main

import (
	"hello/项目的代码结构/DB"
	"hello/项目的代码结构/Model"
	"hello/项目的代码结构/routers"
)

func main() {
	//数据库操作
	err := DB.InitMysql()
	if err != nil {
		panic(err)
	}

	DB.DB.AutoMigrate(&Model.Todo{})

	r := routers.SetupRouter()
	r.Run("8080")

}
