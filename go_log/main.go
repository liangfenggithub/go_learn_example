package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	// logFile, err := os.OpenFile("./run.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//打开文件
	logFile, err := os.OpenFile("./run.log", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	/*
		    log包可以通过SetOutput()方法指定日志输出的方式（Writer），
			但是只能指定一个输出的方式（Writer）。
			我们利用io.MultiWriter()将多个Writer拼成一个Writer使用的特性，
			把log.Println()输出的内容分流到控制台和文件当中。

		    如果只是输出到文件，不输出到终端的话，直接 	log.SetOutput(logFile)就可以了
	*/
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetFlags(log.Lmicroseconds | log.Ldate)
}
func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
