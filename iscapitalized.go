package main

import "fmt"

func IsCapitalized(s string) bool {

	if s == "" {
		return false
	}

	startWord := true

	for _, ch := range s {

		if ch == ' ' {
			startWord = true
			continue
		}

		if startWord {
			if ch >= 'a' && ch <= 'z' {
				return false
			}
			startWord = false
		}
	}
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}