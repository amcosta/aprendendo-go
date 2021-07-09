package fahrenheit

import "dojo-temp-converter/util"

type ConvertToC interface {
	ConvertFToC(temp float64) string
}

type ConvertToK interface {
	ConvertFToK(temp float64) string
}

func ConvertFToC(temp float64) string {
	result := (temp - 32) / 1.8
	return util.FormatCelsius(result)
}

func ConvertFToK(temp float64) string {
	result := (temp-32)/1.8 + 273.15
	return util.FormatKelvin(result)
}
