package main

import "fmt"

func main() {

	personDetails := make(map[string] string)

	personDetails["Jack"] = "Blue"

	fmt.Println(personDetails["Jack"])
}