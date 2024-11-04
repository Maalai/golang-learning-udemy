package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func main() {

	accountBalance, err := fileops.GetFloatFromFile(fileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
	}

	fmt.Println("Welcome to Go Bank!!")
	fmt.Println("Reach us at - ", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is: %v\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit amount is: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount should not be less than 0!!!")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Your new account balance is: %v\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, fileName)
		case 3:
			var withdrawAmount float64
			fmt.Print("Your withdraw amount is: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdraw amount should not be less than 0!!!")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Withdraw amount should not be greather than available balance!!!")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("Your new account balance is: %v\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, fileName)
		default:
			fmt.Println("Thanks for choosing our bank!!")
			fmt.Println("Goodbye!!!")
			return
		}
	}
}
