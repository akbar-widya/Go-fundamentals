package main

import "fmt"

func main() {
	score := 80
	pointerToScore := &score

	fmt.Println("Current store:", *pointerToScore)

	*pointerToScore = 200

	fmt.Println("Updated score:", score, "(var score stores the real value )")
	fmt.Println("Updated score:", *pointerToScore, "(var pointerToScore doesn't store a value, it's store a memory address of 'score'; it's connected to var store)")
}
