package main

import (
	"fmt"
	"strings"
)

func main() {
	var availabe = 50
	var bookingSlice []string
	for {
		var firstName string
		var lastName string
		var email string
		var ticket int

		fmt.Printf("Please enter your FirstName\n")
		fmt.Scan(&firstName)
		fmt.Printf("Please Enter your last name\n")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email\n")
		fmt.Scan(&email)
		fmt.Printf("Enter your tickets number\n")
		fmt.Scan(&ticket)

		if ticket > availabe {
			fmt.Printf("Sorry the tickets have been sold out.\n")
			continue

		}

		bookingSlice = append(bookingSlice, firstName+" "+lastName)

		fmt.Printf("Thank you Mr %v %v for booking %v tickets. You will receive Notification at %v\n", firstName, lastName, ticket, email)

		fmt.Printf("The whole array is %v\n", bookingSlice)

		availabe = availabe - ticket

		fmt.Printf("Availabke Ticket is %v\n", availabe)

		if availabe == 0 {
			fmt.Printf("All tickets have been sold out")
			break
		}

		//for each loop in go syntax is for index, name:= range ArrayName{var something = strings.Fields(name)} //_ is used as blank identifier
		var firstNames []string
		for _, booking := range bookingSlice {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First name of attendes is %v",firstNames)

	}
}
