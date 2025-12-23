package main

import "fmt"

func main() {
	var conference = "Go Conference"
	var remainingTickets = 100

	fmt.Printf("Welcome to the %v. We have %v tickets availabe. Hurry up!!\n", conference, remainingTickets)

	//pointer in go can be defined as variable that stores the memory address of another variable
	var firstName string
	var lastName string
	var email string
	var tickets int
	fmt.Printf("Enter your First Name\n")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your Last Name\n")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email\n")
	fmt.Scan(&email)
	fmt.Printf("Enter the number of tickets you want to book\n")
	fmt.Scan(&tickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirm mail at %v after payment and 30 minutes before program starts\n", firstName, lastName, tickets, email)

	//ticket booking logic
	remainingTickets -= tickets
	fmt.Printf("The remaining tickets are %v\n", remainingTickets)

}
