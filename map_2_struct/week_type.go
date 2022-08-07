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

func main() {
	m := map[string]interface{}{
		"name":   123,
		"age":    "18",
		"emails": []int{1, 2, 3},
	}

	var p Person
	// 对结构体字段类型和map[string]interface{}的对应键值不做强类型一致的校验,
	// 可以使用WeakDecode/WeakDecodeMetadata方法，它们会尝试做类型转换：
	err := mapstructure.WeakDecode(m, &p)
	if err == nil {
		fmt.Printf("person:%#v", p)
	} else {
		fmt.Println(err.Error())
	}

	//结果  person:main.Person{Name:"123", Age:18, Emails:[]string{"1", "2", "3"}} 他将json字符串中类型转换成对应的结构体中的类型
}
