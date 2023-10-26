package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 用os包读取数据
func ReadFromFile() {
	//打开文件
	file, err := os.Open("a.txt") //当前目录 -- 相对路径
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//读取文件
		var temp = make([]byte, 128)
		n, err := file.Read(temp) //这个函数返回 n是读取的字节数，err,和temp这个是直接从文件中读取改变了temp中的值
		if err == io.EOF {        //end of file
			//表示读到了文件的末尾
			fmt.Println(temp[:n]) //输出剩余的字节
		}
		if err != nil {
			fmt.Printf("错误是%v", err)
		}
		fmt.Printf("read %d bytes form files", n)
		fmt.Println(string(temp[:n])) //temp是切片类型的， 所以要转换成字符串
	}
	//关闭文件
	defer file.Close()
}

// 用bufio 读取数据   buf是缓冲区的意思
func ReadByBufio() {
	file, err := os.Open("a.txt") //当前目录 -- 相对路径
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n') //delim 分隔点的意思， 意思是读到哪里
	if err != nil {
		fmt.Println("read by bufio is wrong")
		return
	}
	fmt.Println(line)
	//最后可以通过循环 遍历所有
}

// ioutil 来读取文件  -- 不要用它读大段的文本
//ioutil被弃用了 太可惜了........

// 写文件  os.openfile
func WriteIntoFile() {
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("openfile wrong", err)
		return
	}
	defer file.Close()
	str := "沙河小王子"
	file.Write([]byte(str))
	file.WriteString(str)
}

// 用 bufio.writer来写文件
func WriteByBufio() {
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("openfile wrong", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("沙赫纳扎") //是写在缓存区中，并没有落到磁盘中
	}
	writer.Flush() //是吧缓冲区中的内容存到磁盘中
}

func main() {
	ReadFromFile()
	ReadByBufio()
	WriteIntoFile()
	WriteByBufio()
}

//文件权限  0644
/*
os.O_WRONLY -- 只写
os.O_RDONLY -- 只读
os.O_CREATE -- 新建
os.O_RDWR --读写
os.O_TRUNC --清空
os.O_APPEND --增加
*/
