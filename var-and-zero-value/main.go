package main

import "fmt"

var (
	appName string = "GoLearner"
	version int = 2
)

const Pi = 3.14

func main() {
	fmt.Println("App:", appName, "v", version)
	fmt.Println("Pi is roughly", Pi)
}
