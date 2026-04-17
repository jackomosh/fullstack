package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, ch := range arg {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}


func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
