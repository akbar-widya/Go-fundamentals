package main

import "fmt"

func main() {
	age := 25

	// Use the & operator to grab a variable's address, 
	// and the * operator to read or update the data stored at that specific location.
	var agePointer *int = &age

	fmt.Println("Value of age:", age)
	fmt.Println("Memory address of age:", agePointer)
}