package main

import "fmt"

// You can share the return type just like regular parameters
func divideWithRemainder(a, b int) (quotient, remainder int) {
	if b ==0 {
		return 0, 0
	}

	quotient = a / b	// named return type so doesn't need to do short declaration
	remainder = a % b

	return
}

func main() {
	q, r := divideWithRemainder(240, 7)
	fmt.Println("Quotient:", q, "Remainder:", r)
}