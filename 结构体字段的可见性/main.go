package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//结构体字段的可见性与JSON序列化

//如果一个go语言的包定义的标识符是大写的，则是对外可见的，小写是包内的，但是json是其他包的，所以json 也是不可见的

// JSON 是 前后端数据交互 的 统一轻量级语言
type Student struct { //tag用反引号包裹的一段内容,类似于表示的名称一样
	ID   int    `json:"id" db:"student..."` //有多个键值对的话，用一个空格来分隔，容错率很差，所以一定只有一个空格
	Name string `json:"ccstudyhyc"`         //其他类型（db是数据库）就 空格分隔， 名称用双引号括起来
	//下面的ID也需要大写，不需要用i
}
type Class struct {
	Title   string
	Student []Student
}

func Newstudent(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}
func main() {
	//创建一个班级变量
	c1 := Class{
		Title:   "火箭101",
		Student: make([]Student, 0, 20)}
	//往班级中添加一个学生
	for i := 0; i < 10; i++ {
		//fmt.sprintf() 是一种格式化字符串的手段， 可以返回一个字符串
		tmpStu := Newstudent(i, fmt.Sprintf("stu%02d", i))
		//有一种 包叫 strconv  它可以实现字符串与其他类型的转换
		c1.Student = append(c1.Student, tmpStu)

	}
	//JSON就是把 fun 语言的数据变成 JSON中的数据
	data, err := json.Marshal(c1)
	if err != nil {
		log.Println("mashal is wrong")
	}
	fmt.Printf("json下的数据类型：%T\n", data)
	fmt.Printf("json下的数据展示：%s\n", data)
	// []byte 切片 如果用byte 输出，则会把相应的结构体啊，字符串啊，数字啊都转化为Unicode的数字
	//如果是用%v 输出将会是一团数字
	//要用%s 进行输出

	//Json 的反序列化
	var c2 Class
	json.Unmarshal([]byte(data), &c2)
	//unmarshal 需要前面是一个 []byte类型， 后面要修改c2，所以要传入一个c2 c2也是载体
	fmt.Printf("反json下的数据类型：%T\n", c2)
	fmt.Printf("反json下的数据展示：%v\n", c2)

	//fmt.Println(c1)
}

/*下面开始讨论 []byte 与其他切片的不同
按照某个人（我）的观点， 字节切片完全就可以代替 其他类型的切片，
不代替的原因就是 字节切片是从底层出发的，转化为二进制的
而其他类型的切片是 更加抽象的，将字节按照特定的规律排布的，
但是更加抽象的数据往往更加复杂难以维护。。
*/
