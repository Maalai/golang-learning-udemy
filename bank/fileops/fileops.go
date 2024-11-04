package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	valueByteArray, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("file not found")
	}

	valueString := string(valueByteArray)
	value, err := strconv.ParseFloat(valueString, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored float value in file")
	}
	return value, nil
}
