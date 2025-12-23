package main

import "fmt"

//learning about data types in Go
func main() {
	/*
		var userName //warning because data type is not defined
		userName = "Kiran"
	*/
	//string data type
	var userName string
	userName = "Kiran"
	fmt.Printf("Welcome to the course %v\n", userName)

	//int data type
	var userAge int
	userAge = 20
	fmt.Printf("%v is %v years old\n", userName, userAge)

	//printing the data type 
	var conference = "Go Conference"
	var tickets = 100
	var remainingTickets = 75
	var ticketPrice = 299.99
	
	fmt.Printf("Data type of conference is %T , tickets is %T, remainingTickets is %T, ticketPrice is %T\n", conference, tickets, remainingTickets, ticketPrice)
}
