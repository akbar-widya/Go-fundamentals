package main

import "fmt"

func main() {
	var retryCount int
	var exactTemperature float64
	var userStatus string
	var isVerified bool

	fmt.Printf("Count: %d\n", retryCount)
	fmt.Printf("Temp: %f\n", exactTemperature)
	fmt.Printf("Status: '%s'\n", userStatus)
	fmt.Printf("Verified: %t\n", isVerified)
}
