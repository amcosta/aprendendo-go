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

	converter, err := factory(originTemp)
	if err != nil {
		panic(err)
	}

	switch targetTemp {
	case "C":
		result := converter.ConvertToC(sourceTemp)
		printOutput(result, targetTemp)
	case "F":
		result := converter.ConvertToF(sourceTemp)
		printOutput(result, targetTemp)
	case "K":
		result := converter.ConvertToK(sourceTemp)
		printOutput(result, targetTemp)
	}

	fmt.Println("Não é possível fazer a conversão de temperaturas")
	os.Exit(1)
}

func factory(originUnit string) (util.Converter, error) {
	switch originUnit {
	case "C":
		return new(celsius.Celsius), nil
	case "F":
		return new(fahrenheit.Fahrenheit), nil
	case "K":
		return new(kelvin.Kelvin), nil
	default:
		return nil, fmt.Errorf("nenhum conversor encontrado para o tipo %s", originUnit)
	}
}

func printOutput(result string, targetTemp string) {
	fmt.Printf("Resultado: %s%s\n", result, targetTemp)
	os.Exit(0)
}
