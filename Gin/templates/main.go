package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// 框架是别人给你搭的一个舞台，然后我们是在舞台中用自己的才艺表现
// 所以用什么框架都无所谓，才艺自身的能力才是更加重要的

func sayHello(w http.ResponseWriter, r *http.Request) { //前者时响应写入器，后者是请求对象  （参数非常重要）
	//fileb, err4 := ioutil.ReadFile("Gin/hello.tet")
	//再fmt.fprintf(fileb)就可以直接输出文件中的所有内容
	var buf = make([]byte, 128)
	file, err2 := os.OpenFile("Gin/hello.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err2 != nil {
		log.Println("openfile was wrong ,error:", err2)
		return
	}
	defer file.Close()
	_, err3 := file.Read(buf)
	if err3 != nil {
		log.Println("file.read was wrong,err:", err3)
		return
	}
	_, err := fmt.Fprintln(w, string(buf)) // 往响应中写入内容  (主要的响应语句)
	if err != nil {
		log.Println("写入响应时出错:", err)
		return
	}
}

func main() {
	http.HandleFunc("/demo", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListenAndServe错误:", err)
		return
	}
}
