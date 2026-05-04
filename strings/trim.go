package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "	Hello, This is my first trim in go "
	fmt.Println(str)
	fmt.Println(strings.TrimRight(str, " "))
	fmt.Println(str)
	fmt.Println(strings.TrimLeft(str, " "))
	fmt.Println(str)
	fmt.Println(strings.Trim(str, " "))
}
