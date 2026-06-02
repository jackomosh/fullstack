package main

import "fmt"

type Students struct {
	firstName string
	lastName string
	grade int
	subjects []string
}

func main() {
	studentOne := Students {
		firstName : "Jack",
		lastName : "Omondi",
		grade : 99,
		subjects : []string {
			"Maths",
			"English",
			"Kiswahili",
		},
	}
	fmt.Println(studentOne)
}