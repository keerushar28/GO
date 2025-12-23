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
	fmt.Printf("%v is %v years old", userName, userAge)
}
