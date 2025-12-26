package main

import "fmt"

func main() {
	fmt.Printf("Learning Arrays and slicess\n")
	var conference = "Go Conference"
	var remainingTickets = 100
	var bookings [60]string

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
	bookings[0] = firstName + " " + lastName
	bookings[10] = "Rahul Sharma"
	bookings[20] = "Rahul Sharma"
	bookings[30] = "Rahul Sharma"

	//printing array information

	fmt.Printf("Whole Array: %v\n", bookings)
	fmt.Printf("First Element: %v\n", bookings[0])
	fmt.Printf("Array Length: %v\n", len(bookings))
	fmt.Printf("Type of array is %T\n", bookings)
	fmt.Printf("The remaining tickets are %v\n", remainingTickets)

}
