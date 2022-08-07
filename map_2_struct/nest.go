package main

// 原文链接：https://blog.csdn.net/darjun/article/details/107703571

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

// 如果源数据中有未映射的值（即结构体中无对应的字段），mapstructure默认会忽略它。
// 我们可以在结构体中定义一个字段，为其设置mapstructure:",remain"标签。这样未映射的值就会添加到这个字段中。注意，这个字段的类型只能为map[string]interface{}或map[interface{}]interface{}。
type Person struct {
	Name  string
	Age   int
	Job   string
	Other map[string]interface{} `mapstructure:",remain"`
}

type Friend1 struct {
	Person
}

//设置mapstructure:",squash"的含义是编码时把该结构体的字段提到父结构中，解码时从父结构再提到子结构中
// 另外需要注意一点，如果父结构体中有同名的字段，那么mapstructure会将JSON 中对应的值同时设置到这两个字段中，即这两个字段有相同的值。
type Friend2 struct {
	Person `mapstructure:",squash"`
}

func main() {
	datas := []string{`
	  { 
		"type": "friend1",
		"person": {
		  "name":"dj",
		  "age":18,
		  "job":"programmer",
		  "height":"1.8m",
		  "handsome": true
		}
	  }
	`,
		`
	  {
		"type": "friend2",
		"name": "dj2",
		"age":18,
		"job":"programmer",
		"height":"1.8m",
		"handsome": true
	  }
	`,
	}

	//遍历json字符串数组
	for _, data := range datas {
		var m map[string]interface{}
		//将json字符串解码成 map[string]interface{} 结构
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		//map[string]interface{} -> struct
		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Printf("friend1:%+v\n", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Printf("friend2:%+v\n", f2)
		}
	}
}
