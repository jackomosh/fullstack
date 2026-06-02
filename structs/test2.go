package main

import (
	"fmt"
	"time"
)

type Employee struct {
	name string
	number int
	boss *Employee
	hired time.Time
}

func main() {

	e := Employee {

		name : "Jack",
		number : 1,
		boss : nil,
		hired : time.Now().Local(),

	}

	fmt.Printf("%T %+[1]v\n" , e)
}