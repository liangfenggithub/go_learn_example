package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
	b := make([]byte, 1)
	os.Stdin.Read(b)

}

/*
交叉编译命令：
因为 MT7621 没有 fpu 也就是浮点处理器，所以指定使用软浮点，这样 go 会通过其他指令来模拟浮点运算，编译出来的程序里面不会包含浮点计算指令，否则调用不存在的 fpu 会出现 Illegal instruction
对 MT7621 无 FPU 的另外一个处理办法是在编译 openwrt 时打开内核的浮点模拟器，这样不需要指定 GOMIPS=softfloat 编译出来的程序也可以运行。不过我觉得既然都是模拟，在内核模拟跟在这里模拟也差不多，所以对无 FPU 的处理器，直接指定 GOMIPS=softfloat 即可
经查发现 GOARCH 可以为 mips/mipsle 分别对应大小端，这里 MT7621 为小端架构，于是将 GOARCH 换成 mipsle 后解决

linux 平台：
GOOS=linux GOARCH=mipsle GOMIPS=softfloat CGO_ENABLED=0 go build main.go

windows平台
在 git bash中执行
GOOS=linux  GOARCH=mipsle  GOMIPS=softfloat  CGO_ENABLED=0 go build main.go

交叉编译教程https://studygolang.com/articles/14376



//以下都没有效果！！！
SET GOOS=linux
SET GOARCH=mipsle
SET GOMIPS=softfloat
SET CGO_ENABLED=0
go build main.go

SET GOOS=linux SET GOARCH=mipsle SET GOMIPS=softfloat SET CGO_ENABLED=0 go build main.go

SET GOOS=linux  GOARCH=mipsle  GOMIPS=softfloat  CGO_ENABLED=0 go build main.go





GOOS=linux GOARCH=mipsle go build -ldflags "-s -w" main.go


*/
