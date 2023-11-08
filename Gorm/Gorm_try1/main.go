package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 结构体中 ---> 数据表
type What struct {
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

	if err != nil {
		fmt.Println("创建表出错：", err)
		return
	}

	//// 清空数据库中的数据
	//err = db.Exec("TRUNCATE TABLE user_infos").Error
	//if err != nil {
	//	fmt.Println("清空数据表出错：", err)
	//	return
	//}

	// 列出一些值
	users := []What{
		{ID: 4, Name: "zhangsan", Gender: "male", Hobby: "basketball"},
		{ID: 5, Name: "lisi", Gender: "female", Hobby: "soccer"},
		{ID: 6, Name: "wangwu", Gender: "male", Hobby: "swimming"},
	}

	// 创建表，自动迁移（将结构体和数据表进行对应）
	err = db.AutoMigrate(&What{})

	// 将值存储到数据库
	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			fmt.Println("创建记录出错：", result.Error)
			return
		}
	}

	// 创建数据行
	fmt.Println("成功创建记录！")
	//var u What
	////查询  表中的第一条代码 保存到u中
	//db.First(&u)
	//fmt.Println("u:\v", u)
	////更新
	//db.Model(&u).Update("Hobby", "ball")
	////删除
	//db.Delete(&u)
}
