package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

// mapstructure执行转换的过程中不可避免地会产生错误，例如 JSON 中某个键的类型与对应 Go 结构体中的字段类型不一致。Decode/DecodeMetadata会返回这些错误 从错误信息中很容易看出哪里出错了。
func main() {
	m := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}

	var p Person
	err := mapstructure.Decode(m, &p)
	if err != nil {
		fmt.Println(err.Error())
	}
}
