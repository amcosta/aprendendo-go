package main

import (
	"dojo-temp-converter/converter"
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
	targetTemp, err := input.GetTargetUnit()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	originTemp, err := input.GetOriginTempUnit()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sourceTemp, err := input.GetSourceTemp()
	if err != nil {
		panic(err)
	}

	converter, err := converter.CreateConverter(originTemp)
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

func printOutput(result string, targetTemp string) {
	fmt.Printf("Resultado: %s%s\n", result, targetTemp)
	os.Exit(0)
}
