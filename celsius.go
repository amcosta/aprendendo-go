package dojotempconverter

func ConvertCToF(temp float64) string {
	result := 1.8*temp + 32
	return FormatFahrenheit(result)
}

func ConvertCToK(temp float64) string {
	result := temp + 273.15
	return FormatKelvin(result)
}
