package main

import (
	"fmt"
	"time"
)


type User struct {
	FirstName string
	LastName string
	Bod string
	createdAt time.Time
}


func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBod := getUserData("Enter your birth of date: ")

	appUser := User{FirstName: userFirstName, LastName: userLastName, Bod: userBod, createdAt: time.Now()}

	outputUserDetail(&appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scan(&input)
	return input
}

func outputUserDetail(user *User) {
	fmt.Printf("First Name: %s. Last Name: %s. Birth of Date: %s.\nThis is printed on: %v", user.FirstName, user.LastName, user.Bod, user.createdAt)
}