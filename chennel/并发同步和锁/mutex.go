package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x int64
	lock sync.Mutex  // 这是一个类型
	wg sync.WaitGroup  //这也是一个类型  读和写差不多的时候
	rwlock sync.RWMutex  //读写锁，当读远大于写的时候可以用
)
//写锁只能存在一个，不能和其他写锁和读硕一起 ，而读锁可以和读硕一起
func read()  {
	lock.Lock()   //开锁
	//rwlock.RLock()  读锁的锁住
	time.Sleep(time.Millisecond)
	lock.Unlock()  //解锁
	//rwlock.RUnlock()  读锁的解锁
	wg.Done()
}

func write()  {
	lock.Lock()
	//rwlock.lock()   写锁的锁住
	x := x+1
	lock.Unlock()
	//rwlock.unlock () 写锁的解锁
	time.Sleep(time.Millisecond*5)
	fmt.Println(x)
	wg.Done()
}
func main() {
	start := time.Now()
	for i:= 0 ;i<1000;i++ {
		wg.Add(1)
		go read()
	}

	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))  //计算耗时多久
}
