package celsius

import "dojo-temp-converter/util"

func ConvertToF(temp float64) string {
	result := 1.8*temp + 32
	return util.FormatFahrenheit(result)
}

func ConvertToK(temp float64) string {
	result := temp + 273.15
	return util.FormatKelvin(result)
}
