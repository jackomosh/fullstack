package main

import "fmt"

func RetainFirstHalf(str string) string {

	if len(str) == 1 || str == "" {
		return str
	}

	for _, ch := range str {
		if ch == ' ' {
			continue
		}
	}
	half := len(str) / 2
	return str[:half]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}