package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// mapstructure当然也可以将 Go 结构体反向解码为map[string]interface{}。
// 在反向解码时，我们可以为某些字段设置mapstructure:",omitempty"。这样当这些字段为默认值时(即未设置值)，就不会出现在结构的map[string]interface{}中：
type Person struct {
	Name string
	Age  int
	Job  string `mapstructure:",omitempty"`
}

func main() {
	p := &Person{
		Name: "dj",
		Age:  18,
	}

	//struct -> map(string)interface{}
	var m map[string]interface{}
	mapstructure.Decode(p, &m)

	//map[string]interface{} -> json 字符串
	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}
