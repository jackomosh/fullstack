package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
	
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("What is your first name? ")
	fmt.Scan(&firstName)

	fmt.Print("What is your last name? ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets would you like to purchase? ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, kindly find the tickets on your email %v cheers!!!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for this %v\n", remainingTickets, conferenceName)

	fmt.Printf("All the bookings so far are: %v\n", bookings)
}
