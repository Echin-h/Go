// 一个文件一个包
// 包的名字不能和文件夹的名字一样，不能有- 这个符号
// main 包 入口
package main

import "hello/package/bbb"

func main() {
	bbb.Say()
}
