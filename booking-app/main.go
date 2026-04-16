package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
	
	var userName string
	var userTickets int

	userName = "Jack"
	userTickets = 2

	fmt.Printf("User %v has booked %v tickets\n", userName, userTickets)
}
