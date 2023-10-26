package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件就是拦截者，适合处理一些公共的能力， 他在客户端和服务端之间，通过了他进入服务端
// 中间件返回的是gin.Handerfunc类型
// gin 框架默认是有中间件在里面的 所以是使用 gin.default() 如果你不想要，或者自己设置， 直接gin.new()
// 中间件的使用过程中如果涉及到 go routine 则只能传 c.copy()复制的值，因为go routine 的进行过程中会改变值，使c.next（）中是不可控的
func indexHandler(c *gin.Context) {
	fmt.Println("index...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// 定义一个中间件m1,用来统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时器
	start := time.Now()
	//c.Next() //调用后序的处理函数，  就是调用indexHandler
	c.Abort() //阻止调用后序的函数
	cost := time.Since(start)
	fmt.Printf("cost： %v", cost)
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Next()
	fmt.Println("m2 out...")
}

func main() {
	r := gin.Default()

	//r.Use(m1)  全局注册中间件m1 这样子下面的路由中间的m1都不需要写了
	r.GET("demo", m1, m2, indexHandler) //放入这边的函数，需要传入c *gin.context
	r.GET("lala", m1, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hehe": "hah",
		})
	})

	r.Run(":8080")
}

// 如果使用的是 c.next()  继续往下运行
// 顺序是  m1 in...  m2 in...  index    m2 out...  m1 out...

//如果使用的是 c.abort()  阻止往下运行
//顺序是  m1 in...  m1 out...

/*func authMiddleware(ok bool)gin.HandlerFunc  {
	//连接数据库库
	//或者一些其他的准备工作
	return func (c *gin.Context){
		//是否登录的判断
		// if 登录成功
		c.Next()
		// else
		c.Abort()
	}
}*/ //标准案例
