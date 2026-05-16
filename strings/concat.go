package main

import "fmt"

func main() {
	wordOne := "Hello"
	wordTwo := "World!"
	concat := wordOne + " " + wordTwo
	fmt.Println(concat)
	fmt.Println(len(concat))
}
