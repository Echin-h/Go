package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Userinfo struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("没有连接")
		return
	} else {
		fmt.Println("连接数据库成功")
	}
	err3 := db.AutoMigrate(&Userinfo{})
	if err3 != nil {
		return
	}
	users := []Userinfo{
		{Name: "liwen", Age: 18},
		{Name: "jay", Age: 41},
		{Name: "JJ", Age: 78},
	}
	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Println(result.Error)
			return
		}
	}
	//db.Create(&users)
	var u Userinfo
	result := db.Where("name = ?", "liwen").First(&u)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	result = db.Delete(&u)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
}

//users := []UserInfo{
//	{ID: 4, Name: "zhangsan", Gender: "male", Hobby: "basketball"},
//	{ID: 5, Name: "lisi", Gender: "female", Hobby: "soccer"},
//	{ID: 6, Name: "wangwu", Gender: "male", Hobby: "swimming"},
//}
//
//// 创建表，自动迁移（将结构体和数据表进行对应）
//err = db.AutoMigrate(&UserInfo{})
//
//// 将值存储到数据库
//for _, user := range users {
//	result := db.Create(&user)
//	if result.Error != nil {
//		fmt.Println("创建记录出错：", result.Error)
//		return
//	}

//var user User
//db.First(&user) //直接将表中第一个赋值给user
//fmt.Println(user)
