package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	normalUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser, err := user.NewAdmin("maalai@gmail.com", "1234")

	if err != nil {
		fmt.Println(err)
		return
	}

	normalUser.PrintUser()

	normalUser.ClearUser()

	normalUser.PrintUser()

	adminUser.PrintUser()
	adminUser.ClearUser()
	adminUser.PrintUser()

	customTypeTest()

}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)
	return value
}

func customTypeTest() {
	var name str = "Maalai"
	name.log()
}
