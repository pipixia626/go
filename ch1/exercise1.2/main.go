package main

import (
	"fmt"
	"os"
)

//修改echo程序，输出参数的索引和值
func main() {
	for key, arg := range os.Args[0:] {
		fmt.Println(key, arg)
	}
}
