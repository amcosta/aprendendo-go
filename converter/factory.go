package converter

import "fmt"

func CreateConverter(unit string) (Converter, error) {
	switch unit {
	case "C":
		return new(Celsius), nil
	case "F":
		return new(Fahrenheit), nil
	case "K":
		return new(Kelvin), nil
	default:
		return nil, fmt.Errorf("nenhum conversor encontrado para o tipo %s", unit)
	}
}
