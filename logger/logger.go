package main

import (
	"log"
	"os"
)

// logger 分为 print,fatal,panic 系列
/*	1，print,fatal,panic
func main() {
	log.Println("这是一条很普通的日志")
	v := "很普通的"
	log.Printf("这是一条封%s日志\n", v)
	log.Fatalln("这是一条会触发fatal的日志")
	log.Panicln("这是一条会触发panic的日志")
}*/

// logger 会打印每条日志信息的日期，时间，默认输出到系统的标准错误，
//Fatal系列 会写入日志信息后调用 os.exit(1)
// panic系列会在写入日志信息后panic

/*	2.log.Setflags：
不设置setflags的时候，默认会输出日期时间。。
'Ldate'：在日志中显示日期。
'Ltime'：在日志中显示时间。
'Lmicroseconds'：在日志中显示微秒级别的时间。
'Llongfile'：在日志中显示完整的文件路径和行号。
'Lshortfile'：在日志中显示文件名和行号。
'LUTC'：使用UTC时间而不是本地时间。
*/
/*func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("这是一条普通的日志")
}*/

/*
log.prefix 配置日志前缀：
*/
/*func main() {
	log.SetFlags(log.Llongfile)
	log.SetPrefix("[沙赫纳扎]")
	log.Println("这是一个小王子的故事")
}*/

// 3.配置日志输出位置
/*func main() {
	//一开始是前面的东西设置
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	log.SetPrefix("[小王子]")
	logfile, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("打开文件有错误，错误是%v \n", err)
	}
	log.SetOutput(logfile) //日志输出位置  反正就是把日志写到这个文件里面
	log.Println("这是一条很普通的日志")
	log.Println("这是一条很普通的日志。")
}*/

// 4.标准logger
/*func init()  {
	//就是把上面的配置加到这个init函数里面
}*/

// 5.创建logger
func main() {
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime) // 第二个是前缀，第三个是setflags 反正感觉都差不多
	logger.Println("这是自定义的logger记录的日志。")
}
