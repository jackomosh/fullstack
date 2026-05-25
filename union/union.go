package main

import (
	"os"
	"fmt"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Printf("\n")
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	combined := str1 + str2
	newStr := ""

	for _, ch := range combined {
		if string(ch) != str1 {
			newStr += string(ch)
		}
	}

	fmt.Println(newStr)

}