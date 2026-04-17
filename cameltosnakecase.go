package main

import "fmt"

func CamelToSnakeCase(s string) string{

	if s == "" {
		return ""
	}

	runes := []rune(s)
	// length := len(runes)

	for i, ch := range runes {

		if !(ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z') {
			return s
		}

		for i > 0 && ch >= 'A' && ch <= 'Z' && runes[i - 1] >= 'A' && runes[i - 1] <= 'Z' {
			return s
		}

	}

	if len(runes) - 1 >= 'A' && len(runes) - 1 <= 'Z' {
		return s
	}

	answer := ""

	for i, char := range runes {
		if i > 0 && char >= 'A' && char <= 'Z' {
			answer += "_"
		}
		answer += string(char)
	}
	return answer
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}