package main

import (
	"fmt"
)

type Person struct {
	age  int
	name string
	sex  bool
}

func main() {

	var i Person = Person{age: 10, sex: true, name: "zhangsan"}

	fmt.Printf("type:%T\n", i)    //输出类型
	fmt.Printf("value:%v\n", i)   //直接输出值
	fmt.Printf("value+:%+v\n", i) //先输出字段名称 在输出值
	fmt.Printf("value#:%#v\n", i) //先输出结构体名字，再输出结构体（字段类型+字段的值）

	fmt.Println("========interface========")
	var interf interface{} = i
	fmt.Printf("%v\n", interf)
	fmt.Println(interf)
}
