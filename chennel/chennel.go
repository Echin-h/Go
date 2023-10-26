//单纯地将函数并发执行是没有意义的，函数与函数之间需要交换数据才能体现并发执行函数的意义
//chennel 是一种特殊类型,也是一种引用类型，空就是nil  先进先出
//引用类型需要使用make函数来初始化
/*var ch1 chan int
var ch2 chan string
var ch3 chan bool*/
/*ch1 := make(chan int)
ch2 := make(chan string)
ch3 := make(chan bool)*/

//通道操作  ：  发送，接受，关闭
//   <-

/*
x <- 10
x := <- m
<- m
close(ch)  通道会自动关闭，关闭不是必须的
*/
package main

import "fmt"

func main() {
	var ch1 chan int        //一定需要初始化
	ch1 = make(chan int, 1) //后面的数字是表示带缓存区的   带缓冲区
	//	ch1 = make (chan int )   无缓冲区的   同步通道
	ch1 <- 10
	x := <-ch1
	//	len(x)  获取通道中的数量
	//	cap(x)  通道中的容量  这两个基本没啥用
	fmt.Println(x)
	close(ch1)
}

//引用类型
