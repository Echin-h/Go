package main

import (
	"fmt"
	"time"
)

func main() {
	//时间格式化
	now := time.Now()
	time4 := now.Format("2006-01-02-15-04-05")
	fmt.Println(time4)
	time5 := now.Format("2006/01/02 15:04:05")
	fmt.Println(time5)

	//解析字符串格式的时间
	timeStr := "2023/10/21 10:12:20"

	//断定时区
	loc, err := time.LoadLocation("Asia/shanghai")
	if err != nil {
		fmt.Println("error 时区")
	}
	//2.根据时区去解析字符串格式的时间
	nowTime, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc) //输出时间格式 ，现在时间的字符串， 时区
	if err != nil {
		fmt.Println("error 解析字符串")
	}
	fmt.Println(nowTime)
}
