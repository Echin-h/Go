package main

import (
	"html/template"
	"log"
	"net/http"
)

// 遇事不决 ， 先写注释
func SayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	file, err := template.ParseFiles("Gin/HttpTemplates/hello.tmpl")
	if err != nil {
		log.Println("parse err:", err)
		return
	}
	//3.渲染模板
	err1 := file.Execute(w, "小王子") //前者是对象 ，后者是输入信息，也就是渲染内容
	if err1 != nil {
		log.Println("execute err:", err1)
		return
	}
}

func main() {
	http.HandleFunc("/demo", SayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("listenandserve is wrong, err:", err)
		return
	}
}
