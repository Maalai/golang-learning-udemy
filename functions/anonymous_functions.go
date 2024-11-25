package main

import "fmt"

func anonymousMain() {
	numbers := []int{1, 2, 3, 4, 5}

	doubled := transformNumbersAnonymous(&numbers, func(number int) int {
		return number * 2
	})

	tripled := transformNumbersAnonymous(&numbers, func(number int) int {
		return number * 3
	})

	fmt.Println("Anonymous doubled: ", doubled)
	fmt.Println("Anonymous tripled: ", tripled)
}

func transformNumbersAnonymous(numbers *[]int, transform transformFn) []int {
	transformedNumbers := []int{}

	for _, val := range *numbers {
		// Invoking the function passed as an argument
		transformedNumbers = append(transformedNumbers, transform(val))
	}
	return transformedNumbers
}
