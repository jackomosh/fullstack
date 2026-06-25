package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("\n")
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	combined := str1 + str2
	seen := make(map[rune]bool)
	var result []rune
	for _, ch := range combined {
		if !seen[ch] {
			seen[ch] = true
			result = append(result, ch)
		}
	}
	fmt.Println(string(result))
}