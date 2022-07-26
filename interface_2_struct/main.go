package main

// golang将interface{}转换为struct
// 项⽬中需要⽤到golang的队列，container/list，需要放⼊的元素是struct，但是因为golang中list的设计，从list中取出时的类型为interface{}，所以需要想办法把interface{}转换为struct。
// 这⾥需要⽤到interface assertion，具体操作见下⾯代码：
// 运行命令 go run .\interface_2_struct.go
import (
	"container/list"
	"fmt"
	"strconv"
)

type People struct {
	Name string
	Age  int
}

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	l.PushBack(People{"zjw", 1})

	// Iterate through list and print its contents.
	e := l.Front()
	p, ok := (e.Value).(People)
	if ok {
		fmt.Println("Name:" + p.Name)
		fmt.Println("Age:" + strconv.Itoa(p.Age))
	} else {
		fmt.Println("e is not an People")
	}
}
