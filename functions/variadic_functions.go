package main

import (
	"fmt"
)

func variadicFunctionsMain() {
	fmt.Println("Variadic func with 2 args: ", sumNumbers(1, 1, 2))
	fmt.Println("Variadic func with 3 args: ", sumNumbers(1, 1, 2, 3))

	numbers := []int{1, 2, 3, 4}
	fmt.Println("Variadic func with slice expanded: ", sumNumbers(1, numbers...))
}

func sumNumbers(startingVal int, numbers ...int) int {
	sum := startingVal

	for _, val := range numbers {
		sum += val
	}
	return sum
}
