package main

import "fmt"

func LastWord(s string) string{

	end := len(s) - 1

	for end > 0 && s[end] == ' ' {
		end--
	}

	if end == len(s) {
		return "\n"
	}

	start := end

	for start > 0 && s[start] != ' ' {
		start--
	}

	return s[start+1:end+1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}