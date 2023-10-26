package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端的代码
// tcp client  demo
func main() {
	//1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial failed ,err:%v", err)
		return
	}
	defer conn.Close()
	// 2.利用该链接进行数据的发送和接受
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.Trim(s, "\r\n")
		if strings.ToUpper(s) == "Q" {
			return
		}
		//给客户端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed,err:%v\n", err)
			return
		}
		//从服务端接受回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read failed")
			return
		}
		fmt.Printf("收到服务端回复：", string(buf[:n]))
	}
}
