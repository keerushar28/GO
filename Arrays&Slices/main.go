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

	//error dekhauxa because yo array size vanda mathi xa so, ani go le runtime ma haina compile time mai error dekhauxa
	//bookings[62] = "Error out of bounds" //invalid argument: index 62 out of bounds [0:60]compilerInvalidIndex

	//printing array information

	fmt.Printf("Whole Array: %v\n", bookings)
	fmt.Printf("First Element: %v\n", bookings[0])
	fmt.Printf("Array Length: %v\n", len(bookings))
	fmt.Printf("Type of array is %T\n", bookings)
	fmt.Printf("The remaining tickets are %v\n", remainingTickets)

	//problems with arrays is that we need to know the size of the array at the time of declaration
	//and we cannot add or remove elements from the array after it is declared
	//so, we use slices to overcome this problem
	//slice is similar to array but we dont define the size initially, it grows automatically
	//adding elements to slice is only different from array, printing and accessing elements is same

	//lets see a example how to use slices

	var bookingSlice []string
	// add garda slice name = append(Slicename, elements......)
	bookingSlice = append(bookingSlice, "Ram", "Shyam", "Hari", "Sita", "Gita")
	fmt.Printf("The booking slice is %v\n", bookingSlice)
	fmt.Printf("The length of the booking slice is %v\n", len(bookingSlice))
	fmt.Printf("The capacity of the booking slice is %v\n", cap(bookingSlice))
	fmt.Printf("The first element of the booking slice is %v\n", bookingSlice[0])
	fmt.Printf("The last element of the booking slice is %v\n", bookingSlice[len(bookingSlice)-1])
	fmt.Printf("The capacity of the booking slice is %v\n", cap(bookingSlice))

}
