package main

import "fmt"

type nazhar interface { //接口的嵌套
	mover
	sayer
}

type mover interface {
	move()
}
type sayer interface {
	say()
}
type Person struct {
	name string
	age  int8
}

/*func (p Person) move() {
	fmt.Printf("%s在跑....", p.name)
}*/

func (p Person) move() {
	fmt.Printf("%s在跑....\n", p.name)
}
func (p Person) say() {
	fmt.Printf("%s在叫...\n", p.name)
}

func main() {
	var m nazhar
	p1 := Person{name: "小王子", age: 18}
	//	p2 := &Person{name: "娜扎", age: 16}
	m = p1
	//	m = p2
	m.move()
	m.say()
	fmt.Println(m) //输出结构体，因为m类型赋值 了一个p1
}

//使用值接受者：类型的值和类型的指针都能够保存到接口类型中
//使用指针接收者：值不能传到接口中，
//一个类型可以实现多个接口，一个接口也可以有多个类型（对象）
