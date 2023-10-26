package main

import (
	"fmt"
	"time"
)

//时间间隔的概念

func main() {
	/*time.Duration()*/
	time.Sleep(time.Second * 5)
	//time.Sleep(time.Duration(n)*time.Second)

	//时间计时器
	for temp := range time.Tick(time.Second * time.Duration(5)) {
		fmt.Println(temp) //time.tick 返回的是现在的时间戳
	}
}
