package main

import "fmt"

// custom type for a function which accepts int and returns int
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubled := transformNumbers(&numbers, getTransformerFunction("double"))
	tripled := transformNumbers(&numbers, getTransformerFunction("triple"))

	fmt.Println(doubled)
	fmt.Println(tripled)

	// Calling anonymous_functions logic
	anonymousMain()

	// Calling closures logic
	closuresMain()

	// Calling recursion logic
	recursionMain()

	// Calling variadtic_functions logic
	variadicFunctionsMain()
}

// Function which accepts another function as argument
func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformedNumbers := []int{}

	for _, val := range *numbers {
		// Invoking the function passed as an argument
		transformedNumbers = append(transformedNumbers, transform(val))
	}
	return transformedNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(key string) transformFn {
	if key == "double" {
		return double
	} else {
		return triple
	}
}
