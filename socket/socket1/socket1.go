package main

import (
	"bufio"
	"fmt"
	"net"
)

// server 端
//OSI  应用层，表示层，会话层，传输层，网络层，数据链路层，物理层
//socket 是在应用层和传输层之中， 是面向socket层
//TCP/IP 协议  是传输控制协议，是传输层通信协议

func process(conn net.Conn) {
	defer conn.Close()
	//针对当前的链接做信息的发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("error")
			break
		}
		recv := string(buf[:n])
		fmt.Println("收到的数据:", recv)
		conn.Write([]byte(recv)) //把收到的数据返回给客户端
	}
}

// TCP 服务端程序 tcp server demo
func main() {
	//1.开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed ,err:%v\n", err)
		return
	}
	for {
		//2.等待用户来接受连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue

		}
		//3.启动一个单独的goroutine 去处理链接
		go process(conn)
	}

}
