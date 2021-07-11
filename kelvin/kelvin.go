package kelvin

import "dojo-temp-converter/util"

func ConvertToC(temp float64) string {
	result := temp*-1 + 273.15
	return util.FormatCelsius(result)
}

func ConvertToF(temp float64) string {
	result := (temp-273.15)*1.8 + 32
	return util.FormatCelsius(result)
}
