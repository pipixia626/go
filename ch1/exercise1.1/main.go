package main

import (
	"fmt"
	"os"
)

//修改echo程序输出os.Args[0]，即命令的名字
func main() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
