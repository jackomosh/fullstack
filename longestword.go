package main

import (
	"strings"
	"regexp"
	"fmt"
)

func longestWord(str string) string{
	if len(str) == 0 {
		return ""
	}
	var re = regexp.MustCompile("[^a-zA-Z0-9]+")
	words := strings.Fields(str)
	longWord := ""
	for _, word := range words {
		cleaned := re.ReplaceAllString(word, "")

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
