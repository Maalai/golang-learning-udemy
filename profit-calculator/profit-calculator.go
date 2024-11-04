package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Revenue input is invalid. Exiting the app")
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("Expenses input is invalid. Exiting the app")
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("TaxRate input is invalid. Exiting the app")
		return
	}

	ebt, profit, ratio := getResults(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
	ebtString := fmt.Sprintln("EBT is: ", ebt)
	profitString := fmt.Sprintln("Profit is: ", profit)
	ratioString := fmt.Sprintln("Ratio is: ", ratio)

	completeString := fmt.Sprint(ebtString, profitString, ratioString)

	os.WriteFile("profit-details.txt", []byte(completeString), 0644)

}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		errorObj := errors.New("invalid input")
		return 0.0, errorObj
	}
	return userInput, nil
}

func getResults(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
