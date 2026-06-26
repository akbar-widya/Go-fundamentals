package main

import "fmt"

func main() {
    regularInt := 42

    // Approach 1: Declare first, assign later
    var pointerA *int
    pointerA = &regularInt

    // Approach 2: Declare and assign at the same time
    pointerB := &regularInt

    // Both pointers now hold the exact same memory address
    fmt.Println("pointerA holds:", pointerA)
    fmt.Println("pointerB holds:", pointerB)
}