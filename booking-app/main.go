package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
	
	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidUserTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets, kindly find the tickets on your email %v cheers!!!\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for this %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var name = strings.Fields(booking)
				firstNames = append(firstNames, name[0])
			}

			if remainingTickets == 0 {
				fmt.Println("All tickets are booked, come back next month")
				break
			}

			fmt.Printf("All the bookings so far are: %v\n", firstNames)

		} else {

			if !isValidName {
				fmt.Println("Enter correct names: Names must be more than 2 characters")
			}

			if !isValidEmail {
				fmt.Println("Enter correct email address")
			}
			
			if !isValidUserTickets {
				fmt.Println("Sorry!! Available tickets are %v", remainingTickets)
			}
		}
	}
	fmt.Println("Your slot has been booked succefully")
}
