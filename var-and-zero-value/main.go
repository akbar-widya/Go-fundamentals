package main

import "fmt"

func main() {
	var scores []int	// a slice (dynamic list)
	var userPreferences map[string]string		// a map (key-value store)

	fmt.Println("Scores is nil:", scores == nil)
	fmt.Println("Preferences is nil:", userPreferences == nil)
}
