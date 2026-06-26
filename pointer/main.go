package main

import "fmt"

func main() {
	scorePointer := new(int)

	fmt.Println("Initial value:", *scorePointer)

	*scorePointer = 89
	fmt.Println("Updated value:", *scorePointer)
}
