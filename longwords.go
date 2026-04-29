package main

import (
	"strings"
	"unicode"
	"fmt"
)

func longestWord(str string) string{
	if len(str) == 0 {
		return ""
	}
	words := strings.Fields(str)
	longWord := ""
	for _, word := range words {
		cleaned := ""
		for _, ch := range word {
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
				cleaned += string(ch)
			}
		}
		if len(cleaned) > len(longWord) {
			longWord = cleaned
		}
	}
	return longWord
}

func main() {
	test1 := "The Tiger and the Panda!!"
	test2 := "Check this: @apple, #banana, &cherry."
	test3 := "User123!! vs Admin999??"
	fmt.Printf("Test 1 (Equal length): %s\n", longestWord(test1)) 
	fmt.Printf("Test 2 (Punctuation):  %s\n", longestWord(test2))
	fmt.Printf("Test 3 (Numbers):      %s\n", longestWord(test3))
}
