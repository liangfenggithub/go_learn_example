package main

import (
	"fmt"
	"time"
)

const CLASSB_BEACON_INTERVAL uint64 = 128000 //128s
func getLastBeacon() (nowTimeMs uint64, lastBeaconTimeMs uint64) {
	//得到ms级别的当前时间
	nowTimeMs = uint64(time.Now().UnixMilli())

	//得到上一个beacon的时间
	lastBeaconTimeMs = nowTimeMs - (nowTimeMs % CLASSB_BEACON_INTERVAL)

	fmt.Printf("now: %v lastBeacon: %v\n", nowTimeMs, lastBeaconTimeMs)
	return
}
func waitUtilNextSecond() {
	//等待完整一秒开始时发送，防止在不同秒内发送数据
	nowTimeMs, lastBeaconTimeMs := getLastBeacon()
	fmt.Printf("now: %v lastBeacon: %v\n", nowTimeMs, lastBeaconTimeMs)
	//获取下一秒
	nextSecond := ((nowTimeMs + 1000) / 1000) * 1000
	fmt.Printf("nextSecond time is %v\n", nextSecond)

	//计算需要等待多少us得到下一秒
	waitMs := (nextSecond * 1000) - uint64(time.Now().UnixMicro())
	fmt.Printf("waitMs is %v\n", waitMs)
	if waitMs > 10 {
		waitMs = waitMs - 10
	}

	time.Sleep(time.Microsecond * time.Duration(waitMs))
	fmt.Println("now is", time.Now().Local().UnixMilli())

	return
}

func main() {

	//获取当前时间戳
	timeUnix := time.Now().Unix() //单位s,打印结果:1491888244
	fmt.Println("timeUnix s is:", timeUnix)

	timeUnixNano := time.Now().UnixNano() //单位纳秒,打印结果：1491888244752784461
	fmt.Println("timeUnixNano is:", timeUnixNano)

	timeUnixMs := time.Now().Local().UnixMilli()
	fmt.Println("timeUnixMs is：", timeUnixMs)

	timeUnixUs := time.Now().Local().UnixMicro()
	fmt.Println("timeUnixUs is：", timeUnixUs)

	//时间戳转时间字符串
	timeStr := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	fmt.Println(timeStr)                                //打印结果：2017-04-11 13:24:04

}
