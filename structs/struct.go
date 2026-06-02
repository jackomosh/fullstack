package main

import "fmt"

type Person struct {
	name string
	age int
	favColor string
	weight int
}

func main() {

	person1 := Person {
		name : "Jack",
		age : 28,
		favColor : "Blue",
		weight : 65,
	}

	fmt.Println(person1)
	fmt.Println(person1.age)
	fmt.Println()

	person2 := Person {
		name : "James",
		age : 31,
		favColor : "Black",
		weight : 77,
	}

	fmt.Println(person2)
	fmt.Println(person2.weight)


}