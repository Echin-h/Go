package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	file, err := template.ParseFiles("Gin/TemplateGrammar/hello.html")
	if err != nil {
		log.Println("parsefiles is wrong:", err)
		return
	}
	//渲染模板
	u1 := User{ //u1.Name
		Name:   "饭饭",
		Gender: "女",
		Age:    18,
	}
	m1 := make(map[string]interface{})
	m1["芳龄"] = 18
	m1["我是"] = "太漂亮了"
	//interface{} {填充内容}  前面一个连带大括号都是一个类型

	hobby := []string{
		"打篮球",
		"打足球",
		"打架",
	}
	err1 := file.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobby,
	})
	if err1 != nil {
		log.Println("execute is wrong:", err1)
		return
	}
}

func main() {
	http.HandleFunc("/demo", SayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln("listenandserve is wrong: ", err)
		return
	}
}
