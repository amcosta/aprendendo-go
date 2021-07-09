package kelvin

import "dojo-temp-converter/util"

type ConvertToC interface {
	ConvertKToC(temp float64) string
}

type ConvertToF interface {
	ConvertKToF(temp float64) string
}

func ConvertKToC(temp float64) string {
	result := temp*-1 + 273.15
	return util.FormatCelsius(result)
}

func ConvertKToF(temp float64) string {
	result := (temp-273.15)*1.8 + 32
	return util.FormatCelsius(result)
}
