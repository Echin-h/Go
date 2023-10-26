package main

import "fmt"

//1. 了解方法和函数的区别
//2.了解 指针接收者和值接收者的区别
//3. 了解如何给系统内定的类型定对象

type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person { //函数
	return &Person{name: name,
		age: age}
}

// 方法是针对与person 类型的
func (p Person) Dream() { //方法
	fmt.Printf("%v 的梦想是学好go语言 \t", p.name)
}
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// 接受者为值类型的时候，不会改变其值
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}
func main() {
	p1 := NewPerson("沙赫纳扎", int8(18))

	p1.Dream()
	fmt.Println(p1.age)
	p1.SetAge(int8(16))
	fmt.Println(p1.age)
	p1.SetAge2(int8(18))
	fmt.Println(p1.age)
}

//方法定义只能 定义自己定义的类型，而不是系统自己定义的类型
//如果要定义
//type myInt int
//func (i myInt) sayHi (){
//	fmt.Println("hi")
//}
