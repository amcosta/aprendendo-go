package dojotempconverter

import (
	"fmt"
)

func ConvertToF(temp float64) string {
	result := 1.8*temp + 32
	return fmt.Sprintf("%.2f", result)
}

func ConvertToK(temp float64) string {
	return "1"
}
