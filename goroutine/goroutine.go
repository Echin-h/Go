package main

import (
	"fmt"
	"runtime"
	"sync"
)

// fun routine demo
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello 娜扎", i)
	wg.Done() //通知wg 把计数器减一
}

// 小弟做事情和主函数做事情 谁快谁慢不固定  所以输出的顺序可能有问题-- hello来不及打印就结束了
func main() { //main函数开启的时候会开启一个main fun routine去执行main 函数
	runtime.GOMAXPROCS(1) //表示有几个CPU核心参与此次协程
	wg.Add(100)           //计数牌加一   多一个小弟而已
	fmt.Println("hello main")
	for i := 0; i < 100; i++ { //fun routine的顺序是不固定的   go也支持匿名函数的运行
		go hello(i) //开启了一个独立的go routine去执行hello 函数  （小弟）
	}
	//time.Sleep(time.Second) //来解决 小弟做事情太慢无法输出，等等小弟的例子--但是这个sleep太生硬了，不推荐
	wg.Wait() //等待所有小弟干完活之后才收兵   当计数器为零时，才会收兵
}

/*{   总操作：
var wg sync.WaitGroup
wg.Add(1)
wg.Done()
wg.Wait()}*/
