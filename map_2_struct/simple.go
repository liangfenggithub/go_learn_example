package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string `mapstructure:"username"`
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {

	//json字符串数组
	datas := []string{`
    { 
      "type": "person",
      "username":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	//遍历字符串数组
	for _, data := range datas {
		var m map[string]interface{}
		//json字符串解码成 map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		//根据类型执行 map[string]interface{} -> struct
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat", cat)
		}
	}
}

// 默认情况下，mapstructure使用结构体中字段的名称做这个映射，例如我们的结构体有一个Name字段，mapstructure解码时会在map[string]interface{}中查找键名name。注意，这里的name是大小写不敏感的！
// 也可以指定映射的字段名。为了做到这一点，我们需要为字段设置mapstructure标签。例如下面使用username代替上例中的name：
