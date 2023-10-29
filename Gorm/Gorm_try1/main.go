package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 结构体中 ---> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库出错：", err)
		return
	} else {
		fmt.Println("成功连接到数据库！")
	}

	// 创建表，自动迁移（将结构体和数据表进行对应）
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		fmt.Println("创建表出错：", err)
		return
	}

	// 清空数据库中的数据
	err = db.Exec("TRUNCATE TABLE user_infos").Error
	if err != nil {
		fmt.Println("清空数据表出错：", err)
		return
	}

	// 列出一些值
	users := []UserInfo{
		{ID: 1, Name: "张三", Gender: "男", Hobby: "篮球"},
		{ID: 2, Name: "李四", Gender: "女", Hobby: "足球"},
		{ID: 3, Name: "王五", Gender: "男", Hobby: "游泳"},
	}

	// 将值存储到数据库
	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			fmt.Println("创建记录出错：", result.Error)
			return
		}
	}

	fmt.Println("成功创建记录！")
}
