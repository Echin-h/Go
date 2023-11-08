package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

//gorm Model  内置的一个结构体
/*type Model struct {
	ID uint `gorm:"primary_key"`
	CreatAt time.Time    //创建的记录
	UpdateAt time.Time   //更新的记录
	DeleteAt *time.Time  //删除的记录
}
type User struct {
	gorm.Model
	Name string
}*/

// 定义一个模型
type User struct { // 有结构体标记
	gorm.Model   //ID默认为主键  通过结构体嵌套
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号  唯一 并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置num为自增类型
	Address      string  `gorm:"index:addr"`      //给 address 字段创建一个addr索引
	IgnoreMe     int     `gorm:"-"`               //忽略本字段
}

// 2. 表名通常为结构体的名称的复数，如果要修改就这样子：
// 这样子就能把表名改为自己想要的名字  desc xiaowangzi;
type Animal struct {
	Name string `gorm:"column:ok_name"` //指定列名
	Age  int64
}

func (Animal) TableName() string {
	return "xiaowangzi"
}

// 禁用复数形式： db.SingularTable(true)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	} else {
		println("成功连接到数据库")
	}
	err1 := db.AutoMigrate(&User{})
	if err1 != nil {
		log.Println(err1)
		return
	} else {
		fmt.Println("生成表成功")
	}

	// 主键 ，表名，列名的约定

	// 1. 修改默认主键  `gorm:"primary_key"`

	db.AutoMigrate(&Animal{})
	//3. 使用user结构体创建一个名字叫 ShaHe 的表
	db.Table("ShaHe").Create(&User{})

}
