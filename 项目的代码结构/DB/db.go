package DB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	_, Err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err != nil {
		return Err
	} else {
		return Err
	}
}
