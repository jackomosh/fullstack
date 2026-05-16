package main

import "fmt"

func main() {
	// var balances [5]int

	var cities [3]string

	cities[0] = "Nairobi"
	cities[1] = "Tokyo"

	fmt.Println(cities[0])

	fmt.Println(len(cities))
}