package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Args []string
}

func validateTempUnit(unit string) bool {
	return unit != "F" && unit != "K" && unit != "C"
}

func (input Input) GetTargetUnit() string {
	targetTemp := strings.ToUpper(input.Args[2])
	if validateTempUnit(targetTemp) {
		fmt.Println("A media de temperatura do segundo argumento esta inválido, os valores são (F, K, C)")
		os.Exit(1)
	}

	return targetTemp
}

func (input Input) GetOriginTempUnit() string {
	arg1 := input.Args[1]
	originTemp := strings.ToUpper(string(arg1[len(arg1)-1]))
	if validateTempUnit(originTemp) {
		fmt.Println("A media de temperatura do primeiro argumento esta inválido, os valores são (F, K, C)")
		os.Exit(1)
	}

	return originTemp
}

func (input Input) GetSourceTemp() float64 {
	sourceTemp, err := strconv.ParseFloat(os.Args[1][:len(os.Args[1])-1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return sourceTemp
}
