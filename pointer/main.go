package main

import "fmt"

type User struct {
	Name string
	Score int
}

func resetScore(u User) {
	u.Score = 0
}

func main() {
	user := User{Name: "alice", Score: 10}

	resetScore(user)

	fmt.Println("Score after reset:", user.Score)
}