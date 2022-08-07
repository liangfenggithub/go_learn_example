package main

import (
	"fmt"
	"log"
)

type student struct {
	Name string
	Age  int
}

//演示 struct map修改内部struct中成员的属性
func main() {

	// 方式1：map存储结构体的值，修改属性时，先将结构体整体复制出来，修改后再放进去

	//新建一个map
	stu_map := make(map[string]student)

	//初始化赋值
	stu_map["zs"] = student{ //字面量赋值方式1
		Name: "zhangshan",
		Age:  12,
	}
	stu_map["ls"] = student{"lisi", 22} //字面量赋值方式2，必须严格按照key申明的属性来初始化

	//打印map
	fmt.Printf("stu_map is:%+v\n", stu_map)

	//修改map中student的属性，
	// 注意：这里不可以直接修改，否则会报错: cannot assign to struct field stu_map["zs"].Age in map
	// stu_map["zs"].Age = 22
	zs, ok := stu_map["zs"]
	if ok {
		zs.Age = 22
		stu_map["zs"] = zs
	}

	fmt.Printf("after edit,stu_map is:%+v\n", stu_map)

	//方式二，直接利用student指针存储，而不是值，从而达到直接修改的目的
	fmt.Println("")

	stu_map2 := make(map[string]*student)

	//初始化赋值
	stu_map2["zs"] = &student{ //字面量赋值方式1
		Name: "zhangshan",
		Age:  12,
	}
	stu_map2["ls"] = &student{"lisi", 22} //字面量赋值方式2，必须严格按照key申明的属性来初始化
	fmt.Printf("stu_map2 is:\n")

	for _, v := range stu_map2 {
		log.Printf("%+v\n", v)
	}

	//修改map中student的属性，可以直接修改
	stu_map2["zs"].Age = 22
	fmt.Printf("after edit, stu_map2 is:\n")

	for _, v := range stu_map2 {
		log.Printf("%+v\n", v)
	}

}
