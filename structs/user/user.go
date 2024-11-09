package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	pwd   string
	User
}

// Receiver function of User struct
func (u User) PrintUser() {
	fmt.Printf("User name is %v %v\n", u.firstName, u.lastName)
	fmt.Println("Birthdate is (DD/MM/YYYY): ", u.birthdate)
}

// Receiver function of User struct with pass by reference
func (u *User) ClearUser() {
	u.birthdate = ""
}

// Utility finction to create a new instance of this User struct
func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}
	return &User{firstName: firstName, lastName: lastName, birthdate: birthdate, createdAt: time.Now()}, nil
}

func NewAdmin(email, pwd string) (*Admin, error) {
	return &Admin{email: email, pwd: pwd, User: User{firstName: "ADMIN", lastName: "ADMIN", birthdate: "-----"}}, nil
}
