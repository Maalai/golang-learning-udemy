package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnValue float64
	var years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Value: ")
	fmt.Scan(&expectedReturnValue)

	outputText("Years of Investment: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnValue, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value(adjusted after Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnValue, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnValue/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
