package main

import (
	"dojo-temp-converter/celsius"
	"dojo-temp-converter/fahrenheit"
	"dojo-temp-converter/kelvin"
	"dojo-temp-converter/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong input, please inform <source> <target>")
	}

	targetTemp := strings.ToUpper(os.Args[2])
	if targetTemp != "F" && targetTemp != "K" && targetTemp != "C" {
		fmt.Println("A media de temperatura do segundo argumento esta inválido, os valores são (F, K, C)")
		os.Exit(1)
	}

	arg1 := os.Args[1]
	originTemp := strings.ToUpper(string(arg1[len(arg1)-1]))
	sourceTemp, err := strconv.ParseFloat(os.Args[1][:len(os.Args[1])-1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if originTemp != "F" && originTemp != "K" && originTemp != "C" {
		fmt.Println("A media de temperatura do primeiro argumento esta inválido, os valores são (F, K, C)")
		os.Exit(1)
	}

	switch originTemp {
	case "C":
		result := convertToCelsius(sourceTemp, targetTemp)
		printOutput(result, targetTemp)
	case "F":
		result := convertToFahrenheit(sourceTemp, targetTemp)
		printOutput(result, targetTemp)
	case "K":
		result := convertToKelvin(sourceTemp, targetTemp)
		printOutput(result, targetTemp)
	}
}

func printOutput(result string, targetTemp string) {
	fmt.Printf("Resultado: %s%s\n", result, targetTemp)
}

func convertToCelsius(sourceTemp float64, targetTemp string) string {
	switch targetTemp {
	case "F":
		return celsius.ConvertToF(sourceTemp)
	case "K":
		return celsius.ConvertToK(sourceTemp)
	default:
		return util.FormatCelsius(sourceTemp)
	}
}

func convertToFahrenheit(sourceTemp float64, targetTemp string) string {
	switch targetTemp {
	case "C":
		return fahrenheit.ConvertToC(sourceTemp)
	case "K":
		return fahrenheit.ConvertToK(sourceTemp)
	default:
		return util.FormatFahrenheit(sourceTemp)
	}
}

func convertToKelvin(sourceTemp float64, targetTemp string) string {
	switch targetTemp {
	case "C":
		return kelvin.ConvertToC(sourceTemp)
	case "F":
		return kelvin.ConvertToF(sourceTemp)
	default:
		return util.FormatKelvin(sourceTemp)
	}
}
