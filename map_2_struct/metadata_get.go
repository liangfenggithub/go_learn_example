package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// 解码时会产生一些有用的信息，mapstructure可以使用Metadata收集这些信息。Metadata结构如下：

// type Metadata struct {
//   Keys   []string
//   Unused []string
// }

// Metadata只有两个导出字段：

// Keys：解码成功的键名；
// Unused：在源数据中存在，但是目标结构中不存在的键名。
// 为了收集这些数据，我们需要使用DecodeMetadata来代替Decode方法

type Person struct {
	Name string
	Age  int
}

func main() {
	m := map[string]interface{}{
		"name": "dj",
		"age":  18,
		"job":  "programmer",
	}

	var p Person
	var metadata mapstructure.Metadata
	//利用方法DecodeMetadata解码得到哪些字段存在且可以解码，那么存在不匹配结构体无法解码
	mapstructure.DecodeMetadata(m, &p, &metadata)

	fmt.Printf("keys:%#v unused:%#v\n", metadata.Keys, metadata.Unused)
}
