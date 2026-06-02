package main

import "fmt"

type Doctor struct {
	number int
	name string
	companions []string
}

func main() {

	aDoctor := Doctor {
		number : 3,
		name : "Jack",
		companions : []string {
			"Charity",
			"Audrey",
			"Riziki",
		},
	}
	fmt.Println(aDoctor.companions[0])
}