package main

import "fmt"

/*
两个go routine
	一个goroutine 生成0-100发送到ch1
	一个goroutine 从ch1中取出数据计算他 的平方，然后把结果发送到ch2
*/

func f1(ch chan int) { //单向通道，一般在函数传参的时候限定只能发送或者接受
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// chan <- int 只能写 不能读
// <- chan int 只能读 不能写

func f2(ch1 chan int, ch2 chan int) {
	//从通道中取值的方式一
	for { //无限循环
		tmp, ok := <-ch1 //将ch1传入tmp中,如果ch1空了，则返回false
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}
func main() {

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	for ret := range ch2 { //一般使用for range 来遍历通道
		fmt.Println(ret)
	} //当这个通道为空的时候，for range 会自动跳出循环
}
