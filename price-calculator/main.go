package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxes := []float64{0, 0.7, 0.1, 0.15}
	for _, tax := range taxes {
		ioManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		priceJob := prices.NewTaxIncludedPriceJob(tax, ioManager)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}
}
