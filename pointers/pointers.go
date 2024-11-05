package main

import "fmt"

func main() {
	age := 35

	fmt.Println(getAdultYears(&age))
	editAgeToAdultYears(&age)
	fmt.Println(age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
