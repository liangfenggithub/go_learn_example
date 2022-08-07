package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间戳
	timeUnix := time.Now().Unix() //单位s,打印结果:1491888244
	fmt.Println("time s is:", timeUnix)
	timeUnixNano := time.Now().UnixNano() //单位纳秒,打印结果：1491888244752784461
	fmt.Println("time ns is:", timeUnixNano)

	fmt.Println("time ms is：%v", timeUnixNano/1e6)

	//时间戳转时间字符串
	timeStr := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	fmt.Println(timeStr)                                //打印结果：2017-04-11 13:24:04

}
