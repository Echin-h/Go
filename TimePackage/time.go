package main

import (
	"fmt"
	"time"
)

// time 包
func main() {
	/*now := time.Now() //时间对象
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)

	//时间戳的概念：  1970年1.1到现在的一个秒数
	now := time.Now()
	timeStamp1 := now.Unix()     //将现在的时间转化为总秒数
	timeStamp2 := now.UnixNano() //纳秒数
	fmt.Println(timeStamp1, timeStamp2)
	// 将时间戳转化为具体的时间格式
	//1697815780
	t := time.Unix(1697815780, 0) //将总秒数转化为具体的时间
	fmt.Println(t)

	//时间间隔  time.duration（）
	time.Sleep(5 * time.Second) //睡眠五秒   5这边的类型是时间间隔类型，而不是单纯的数字类型   second是代表秒
	fmt.Println("over")
	//n := 5
	//time.Sleep(time.Duration(n)*time.Second)*/

	/*now := time.Now()
	fmt.Println(now)
	// now + 1hour  add
	t2 := now.Add(time.Hour)
	fmt.Println(t2)
	//sub
	fmt.Println(t2.Sub(now))*/

	//定时器
	/*for temp := range time.Tick(time.Second) {
		fmt.Println(temp)
	}*/

	//时间格式化
	//2006 01 02 15 04 05
	/*now := time.Now()
	ret := now.Format("2006-01-02-15-04-05") //字符串中的数字一定不能填错，但是‘-’可以随便填
	fmt.Println(ret)
	ret2 := now.Format("2006-01-02 15.04.05")
	fmt.Println(ret2)*/

	//解析字符串格式的时间
	//1.拿到时区
	timeStr := "2023-10-20-23-51-24"
	loc, err := time.LoadLocation("Asia/Shanghai") //字符串里写的是一个时区
	if err != nil {
		fmt.Println("error1")
		return
	}
	//2.根据时区去解析字符串格式的时间
	timeObj2, err := time.ParseInLocation("2006-01-02-15-04-05", timeStr, loc) //表示某个时区的时间  这个时间格式要和上面的时间格式要一致
	/*time.Parse("2006/01/02 15:04:05", timeStr)  */ //表示 指定某个时区的时间
	if err != nil {
		fmt.Println("error2")
		return
	}
	fmt.Println(timeObj2)
}
