// 空接口  空接口是没有定义任何方法的接口，所以任何类型都可以实现空接口
// 即空接口的变量可以蠢猪任意值
//
//	空接口+类型判断：
//
// 空接口一般不需要提前定义
package main

import "fmt"

type xxx interface{}

// 应用： 空接口作为函数的参数
//
//	空接口可以作为map的value
func main() {
	var x interface{}
	x = "hello"
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	x = false
	fmt.Println(x)

	/*var m = make(map[string]interface{}, 16)
	m["name"] = "娜扎"
	m["age"] = int8(9)
	m["hobby"] = []string{"篮球", "足球"}
	fmt.Println(m)*/

	//类型断言：

	//如果猜对了
	ret := x.(bool) //猜对了
	fmt.Println(ret)

	//如果没有猜对
	res, ok := x.(string) //猜错了
	if !ok {
		fmt.Println("不是字符类型")
	} else {
		fmt.Println(res)
	}

	//使用switch 进行类型判断
	switch x.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case int8:
		fmt.Println("int8")
	default:
		fmt.Println("[]string")

	}
}
