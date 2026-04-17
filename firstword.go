package main

import "fmt"

func FirstWord(s string) string {
    
	start := 0

	for start < len(s) && s[start] == ' ' {
		start++
	}

	if start == len(s) {
		return "\n"
	}

	end := start

	for end < len(s) && s[end] != ' ' {
		end++
	}

	return s[start:end] + "\n"



}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
