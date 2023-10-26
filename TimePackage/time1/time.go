package main

//表现具体的年月日
import (
	"fmt"
	"time"
)

// 时间加法
func Add() {
	now := time.Now()
	n := 7
	// time.hour 是一个小时单位  后面乘一个时间间隔
	p1 := now.Add(time.Hour * time.Duration(n)) //里面放一个duration
	fmt.Println(p1)
}

// 时间减法
func sub() {
	now := time.Now()
	duration := now.Sub(now.Add(time.Hour * time.Duration(7)))
	fmt.Println("时间差值:", duration)
}

func main() {
	now := time.Now()
	day := now.Day()
	month := now.Month()
	year := now.Year()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
}
