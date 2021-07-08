package main

import "fmt"

func main() {
	s := "hello,world"
	fmt.Println(len(s))
	fmt.Println(s[7:])
	fmt.Println(s[:5])

	b := []byte(s)
	s2 := string(b)
	print(s2)
	fmt.Println(s2)

}
