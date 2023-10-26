package main

import (
	"fmt"
	"time"
)

// 时间转化
// 时间戳的概念：  1970年1.1到现在的一个秒数
func main() {
	//将时间转化为具体的秒数
	now := time.Now()   //获取现在的时间戳
	temp1 := now.Unix() //获得总秒数
	temp2 := now.UnixNano()
	fmt.Println(temp1, temp2)

	//将总秒数转化为具体的时间
	time2 := time.Unix(1697815780, 0) //unix 表示的是 unix时间戳
	fmt.Println(time2)

}
