package main

import "fmt"

func main() {
	ch := make(chan int, 100)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: //如果ch中有值可以传出去，就会被实现
			fmt.Println(x)
		case ch <- 1: //如果ch 有空位可以接受，就会实现
		default: //其他
			fmt.Println("啥都不干")

		}
	}
}
