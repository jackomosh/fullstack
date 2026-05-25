package main

import "fmt"

func main() {

	studentAge := make(map[string] int)
	studentAge["Jack"]	= 28
	studentAge["Hike"] = 22
	studentAge["Leon"] = 21
	studentAge["Olivia"] = 12
	fmt.Println(studentAge)
}
