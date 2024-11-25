package main

import "fmt"

func closuresMain() {
	numbers := []int{1, 2, 3, 4, 5}

	// Returned function from doubled and tripled have their own factors. Since the scope of the factor is locked in as part of closures during it's invocation
	doubled := transformNumbersClosures(&numbers, createTransformerFunction(2))

	tripled := transformNumbersClosures(&numbers, createTransformerFunction(3))

	fmt.Println("Closures doubled: ", doubled)
	fmt.Println("Closures tripled: ", tripled)
}

func transformNumbersClosures(numbers *[]int, transform transformFn) []int {
	transformedNumbers := []int{}

	for _, val := range *numbers {
		// Invoking the function passed as an argument
		transformedNumbers = append(transformedNumbers, transform(val))
	}
	return transformedNumbers
}

func createTransformerFunction(factor int) func(int) int {
	return func(number int) int {
		// factor variable is accesible though the anonymous function doesn't have this parameter. It is passed from parent scope
		return number * factor
	}
}
