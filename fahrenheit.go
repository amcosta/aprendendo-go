package dojotempconverter

func ConvertFToC(temp float64) string {
	result := (temp - 32) / 1.8
	return FormatCelsius(result)
}

func ConvertFToK(temp float64) string {
	result := (temp-32)/1.8 + 273.15
	return FormatKelvin(result)
}
