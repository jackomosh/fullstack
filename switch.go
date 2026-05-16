package main

import "fmt"

func main() {
	age := 18
	switch {
	case age < 18:
		fmt.Println("You cannot vote")
	case age >= 18:
		fmt.Println("You can vote")
	}
}
