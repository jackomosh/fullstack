package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]

	oldChar := os.Args[2]
	newChar := os.Args[3]

	res := ""

	for _, ch := range str {
		if string(ch) == oldChar {
			res += newChar
		} else {
		res += string(ch)
		}
	}
	fmt.Println(res)
}