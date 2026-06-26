package main

import "fmt"

type User struct {
	Name  string
	Score int
}

func incrementScore(u *User) {
	u.Score++
}

func main() {
	user := User{Name: "alice", Score: 10}
	fmt.Println("Score before:", user.Score)

	incrementScore(&user)

	fmt.Println("Score after increment:", user.Score)
}
