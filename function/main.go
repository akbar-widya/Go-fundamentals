package main

import "fmt"

func greet(firstName, lastName string) {
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func main() {
	greet("Alice", "Smith")
}