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

// Receiver function of User struct
func (u User) printUser() {
	fmt.Printf("User name is %v %v\n", u.firstName, u.lastName)
	fmt.Println("Birthdate is (DD/MM/YYYY): ", u.birthdate)
}

// Receiver function of User struct with pass by reference
func (u *User) clearUser() {
	u.birthdate = ""
}

// Utility finction to create a new instance of this User struct
func newUser(firstName, lastName, birthdate string) *User {
	return &User{firstName: firstName, lastName: lastName, birthdate: birthdate, createdAt: time.Now()}
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	user := newUser(firstName, lastName, birthdate)

	user.printUser()

	user.clearUser()

	user.printUser()

}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}
