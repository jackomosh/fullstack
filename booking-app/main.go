package main

import (
	"fmt"
	"booking-app/helper"
	"time"
	"sync"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greatUser()
	
	for {
		
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)


		if isValidName && isValidEmail && isValidUserTickets {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			

			firstNames := printFirstNames()
			fmt.Printf("All first names for the booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are booked, come back next month")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("Enter correct names: Names must be more than 2 characters")
			}

			if !isValidEmail {
				fmt.Println("Enter correct email address")
			}
			
			if !isValidUserTickets {
				fmt.Printf("Sorry!! Available tickets are %v\n", remainingTickets)
			}
		}
	}
	wg.Wait()
	fmt.Println("Your slot has been booked succefully")
}


func greatUser() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)


	fmt.Printf("Thank you %v %v for booking %v tickets, kindly find the tickets on your email %v cheers!!!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for this %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets to %v %v\n", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending tickets.......:\n%v \nTo email %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}