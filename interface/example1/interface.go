package main

import "fmt"

type dog struct{}
type cat struct{}

func (d dog) say() {
	fmt.Println("wangwangwang...")
}
func (c cat) say() {
	fmt.Println("miaomiaomiao...")
}

type sayer interface {
	say()
}

/*
	func do(org sayer) {
		org.say()
	}
*/
func main() {
	var x sayer
	c1 := cat{}
	x = c1
	x.say()
}
