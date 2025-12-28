package main

import (
	"fmt"
	"strings"
)

func main() {
	var remainingTickets = 50
	var bookings []string
	for {

		if remainingTickets == 0 {
			fmt.Printf("See You next year\n")
			break

		}
		var firstName string
		var lastName string
		var email string
		var ticket int
		fmt.Printf("Please enter your first name\n")
		fmt.Scan(&firstName)
		fmt.Printf("Please enter your last name\n")
		fmt.Scan(&lastName)
		fmt.Printf("Please enter your email \n")
		fmt.Scan(&email)
		fmt.Printf("Please enter your ticket\n")
		fmt.Scan(&ticket)

		if remainingTickets >= ticket {
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you for booking with us Mr %v %v you will receive notice at %v\n", firstName, lastName, email)
			remainingTickets = remainingTickets - ticket
		}
		fmt.Printf("Remaining tickets is %v\n", remainingTickets)
		fmt.Printf("The booking array is %v\n",bookings)

		var firstNames []string
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First name is %v\n", firstNames)

	}
}
