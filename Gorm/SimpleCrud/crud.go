package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("无法连接到数据库")
		return
	} else {
		fmt.Println("成功连接到数据库")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println("无法创建表")
		return
	}

	// 创建数据
	user := User{Name: "Alice", Age: 25}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	fmt.Println("创建数据成功:", user)

	// 读取数据
	var u User
	result = db.First(&u, user.ID)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	fmt.Println("读取数据成功:", u)

	// 更新数据
	u.Name = "Bob"
	result = db.Save(&u)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	fmt.Println("更新数据成功:", u)

	// 读取更新后的数据
	var updatedUser User
	result = db.First(&updatedUser, u.ID)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	fmt.Println("更新后的数据:", updatedUser)

	// 删除数据
	result = db.Delete(&u)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	fmt.Println("删除数据成功:", u)
}
