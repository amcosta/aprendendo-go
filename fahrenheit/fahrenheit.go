package fahrenheit

import "dojo-temp-converter/util"

func ConvertToC(temp float64) string {
	result := (temp - 32) / 1.8
	return util.FormatCelsius(result)
}

func ConvertToK(temp float64) string {
	result := (temp-32)/1.8 + 273.15
	return util.FormatKelvin(result)
}
