package converters

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floatPrices []float64
	for _, val := range strings {
		floatVal, err := strconv.ParseFloat(val, 64)

		if err != nil {

			return nil, errors.New("converting price from string to float failed")
		}
		floatPrices = append(floatPrices, floatVal)
	}
	return floatPrices, nil
}
