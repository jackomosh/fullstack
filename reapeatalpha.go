package main

import "fmt"

func RepeatAlpha(s string) string {
	
	res := ""
	count := 1

	for _, ch := range s {

		if ch >= 'A' && ch <= 'Z' {
			count = int((ch)) - 'A' + 1
		} else if ch >= 'a' && ch <= 'z' {
			count = int((ch)) - 'a' + 1
		}

		for i := 0; i < (count); i++ {
			res += string(ch)
		}
	}
	return res
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

