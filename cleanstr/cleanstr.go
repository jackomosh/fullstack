package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Expeted: 1 Input")
		return
	}

	str := os.Args[1]
	var words []string
	currentWord := ""

	for _, ch := range str {
		if ch == ' ' || ch == '\t' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(ch)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	if len(words) == 0 {
		fmt.Println()
		return
	}

	for i, word := range words {
		fmt.Print(word)
		if i < len(words) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}