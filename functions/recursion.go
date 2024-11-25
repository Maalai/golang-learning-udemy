package main

import "fmt"

func recursionMain() {
	fmt.Println("Factorial of 3 is: ", getFactorial(3))
	fmt.Println("Factorial of 4 is: ", getFactorial(4))
	fmt.Println("Factorial of 5 is: ", getFactorial(5))
}

func getFactorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * getFactorial(number-1)
}
