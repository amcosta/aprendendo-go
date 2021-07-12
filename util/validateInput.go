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

func (input Input) GetTargetUnit() (string, error) {
	targetTemp := strings.ToUpper(input.Args[2])
	if validateTempUnit(targetTemp) {
		return "", fmt.Errorf("A media de temperatura do segundo argumento esta inválido, os valores são (F, K, C)")
	}

	return targetTemp, nil
}

func (input Input) GetOriginTempUnit() (string, error) {
	arg1 := input.Args[1]
	originTemp := strings.ToUpper(string(arg1[len(arg1)-1]))
	if validateTempUnit(originTemp) {
		return "", fmt.Errorf("A media de temperatura do primeiro argumento esta inválido, os valores são (F, K, C)")
	}

	return originTemp, nil
}

func (input Input) GetSourceTemp() (float64, error) {
	return strconv.ParseFloat(os.Args[1][:len(os.Args[1])-1], 64)
}
