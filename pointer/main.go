package main

import "fmt"

func main() {
    var activeScore *int 

    fmt.Println("Pointer value:", activeScore) 	// Output: <nil>

	if activeScore == nil {
        fmt.Println("No score is currently tracked.")
    }
}
