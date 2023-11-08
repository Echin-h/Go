package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 1. 定义模型
type User struct {
	ID   int64
	Name string `gorm:"default:'ZhuTou'"` //设定一个默认值，如果没有传入值就是显示ZhuTou
	//但是 如果你想要把结构体赋值为空或者零值时，它依旧会显示一个'ZhuTou'，这个需要注意
	//如果想要给字段 赋值零值，则可以
	Age int64
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
	// 2.把模型与数据库中的表连接一起来
	db.AutoMigrate(&User{})

	//3. 创建记录
	u := User{ID: 123, Name: "cc", Age: 18} //在代码层面创建一个对象
	//给字段赋零值的两个方法
	/* 1.使用指针方式将零值传入
	Name *string  `gorm:"...."`
	u :=User{Name:new(string),...}*/
	/* 2.使用 Scanner/Valuer
	Name sql.NullString `gorm:.....`
	u := User{Name: sql.NullString{"",true},Age:18}
	*/

	db.Create(&u) // 可以传指针或者结构体本身

}
