package celsius

import "dojo-temp-converter/util"

type ConvertToF interface {
	ConvertCToF(temp float64) string
}

type ConvertToK interface {
	ConvertCToK(temp float64) string
}

func ConvertCToF(temp float64) string {
	result := 1.8*temp + 32
	return util.FormatFahrenheit(result)
}

func ConvertCToK(temp float64) string {
	result := temp + 273.15
	return util.FormatKelvin(result)
}
