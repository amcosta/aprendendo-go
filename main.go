package main

import (
	"dojo-temp-converter/celsius"
	"dojo-temp-converter/fahrenheit"
	"dojo-temp-converter/kelvin"
	"dojo-temp-converter/util"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Entrada inválida, por favor informe <unidade-origem> <unidade-destino>")
		os.Exit(1)
	}

	input := util.Input{Args: os.Args}
	targetTemp := input.GetTargetUnit()
	originTemp := input.GetOriginTempUnit()
	sourceTemp := input.GetSourceTemp()

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

	fmt.Println("Não é possível fazer a conversão de temperaturas")
	os.Exit(1)
}

func printOutput(result string, targetTemp string) {
	fmt.Printf("Resultado: %s%s\n", result, targetTemp)
	os.Exit(0)
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
