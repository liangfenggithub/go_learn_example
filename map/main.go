package main

import "fmt"

func main() {
	//定义
	map1 := map[string]int{}
	map2 := make(map[string]string)

	//赋值
	map1["first"] = 1
	map1["second"] = 2
	map1["third"] = 3

	map2["a"] = "aaaa"
	map2["b"] = "bbbb"
	map2["c"] = "cccc"

	//遍历输出
	fmt.Println("map1 value is:")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//读取map值
	fmt.Println("map2 的 b 值为", map2["b"])

	//判断一个key是否存在
	str, ok := map2["c"]
	if ok {
		fmt.Println("map2[\"c\"] = ", str)
	}

	//删除一个元素
	fmt.Println("map2 删除一个key前的值是：")
	for k, v := range map2 {
		fmt.Println(k, v)
	}
	delete(map2, "c")
	fmt.Println("map2 删除一个key后的值是：")
	for k, v := range map2 {
		fmt.Println(k, v)
	}

	//新增一个元素
	map2["d"] = "dddd"
	fmt.Println("map2 新增一个key后的值是：")
	for k, v := range map2 {
		fmt.Println(k, v)
	}

}
