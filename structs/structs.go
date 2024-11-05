package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	user := User{firstName: firstName, lastName: lastName, birthdate: birthdate, createdAt: time.Now()}

	outputUser(&user)

}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}

func outputUser(user *User) {
	fmt.Printf("User name is %v %v\n", user.firstName, user.lastName)
	fmt.Println("Birthdate is (DD/MM/YYYY): ", user.birthdate)

}
