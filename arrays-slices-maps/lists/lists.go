package lists

import "fmt"

func main() {
	prices := []float64{9.99, 1.11, 2.00, 3.18}

	fmt.Println(prices[1])

	// Remove from slice
	prices = prices[0:2]
	fmt.Println(prices)

	// append to slice
	prices = append(prices, 2.10)
	fmt.Println(prices)
	// mainPractise()

	// Unpacking list values
	discountedPrices := []float64{100.99, 80.99, 20.59}

	prices = append(prices, discountedPrices...)
	fmt.Println(prices)
}

/* func main() {
	prices := [4]float64{9.99, 1.11, 2.00, 3.18}

	fmt.Println(prices)

	fmt.Println(prices[:])

	filteredPrices := prices[1:]

	anotherFilteredPrices := filteredPrices[:1]

	fmt.Println(anotherFilteredPrices)

	// overriding the value of a slice
	filteredPrices[0] = 7.11
	fmt.Println(prices)

	// len and cap of slice
	fmt.Println(len((anotherFilteredPrices)), cap(anotherFilteredPrices))

	// getting the element out of the slice which has more capacity
	fmt.Println(anotherFilteredPrices[:3])
} */
