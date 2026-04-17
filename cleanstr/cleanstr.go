package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Check if number of arguments is exactly 1 (plus the program name)
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	str := os.Args[1]
	var words []string
	currentWord := ""

	for _, ch := range str {
		// A "word" is delimited by space OR tab
		if ch == ' ' || ch == '\t' {
			// Only append if currentWord isn't empty (avoids multiple spaces)
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(ch)
		}
	}

	// 2. Don't forget the last word if the string didn't end in a space
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// 3. Handle the "no words to display" case
	if len(words) == 0 {
		fmt.Println()
		return
	}

	// 4. Join the words with exactly one space
	for i, word := range words {
		fmt.Print(word)
		if i < len(words)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println() // Final newline
}