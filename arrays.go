package main

import "fmt"

func main() {
	var EvenNum = [5]int{5, 4, 3, 2, 1}
	for i, value := range EvenNum {
		fmt.Println(value, i)
	}
	fmt.Println(EvenNum[:3])
}
