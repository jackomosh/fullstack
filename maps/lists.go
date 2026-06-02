package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int {
		"Jack":28,
		"Hike":23,
		"Leon":21,
		"James":30,
	}
	fmt.Println(mp["Jack"])
	fmt.Println(mp)
}