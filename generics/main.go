package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.2))
	fmt.Println(add("Maalai", "S"))
}

// Generic function which accepts int or float64 or string as args
func add[T int | float64 | string](a, b T) T {
	return a + b
}
