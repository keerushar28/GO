package main

import "fmt"

func main() {
	fmt.Println(`Hello, World!`)
	fmt.Println("Welcome to Go programming!")
	fmt.Println("This is a simple Go application.")
	fmt.Println("This is demo from IDE from JetBrains ")
	
	// Demonstrate the addition function
	demonstrateAddition()
}

//function to make a addition
func add(a, b int) int {
	return a + b
}

func demonstrateAddition() {
	result := add(5, 3)
	fmt.Printf("Addition of 5 and 3 is: %d\n", result)
}
