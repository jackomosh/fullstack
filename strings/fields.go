package main

import (
	"fmt"
	"strings"
)

func main() {
	answer := " Hello,	This is z01 		\n	Kisumu "
	fmt.Println(strings.Fields(answer))

	input := strings.Fields(answer)
	input1 := strings.Join(input, " ")

	fmt.Println(input1)

	fmt.Println(string(answer[9]))
}
