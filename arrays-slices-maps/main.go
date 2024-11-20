package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {

	userNames := make([]string, 2, 5)

	userNames[0] = "Maalai"
	userNames[1] = "S"

	userNames = append(userNames, "H")

	// fmt.Println(userNames)

	courseRatings := make(floatMap, 2)

	courseRatings["Go"] = 5.0
	courseRatings["React"] = 5.0

	// courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, val := range courseRatings {
		fmt.Println(key, val)
	}

}
