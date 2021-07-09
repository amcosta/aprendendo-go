package dojotempconverter

func ConvertKToC(temp float64) string {
	result := temp*-1 + 273.15
	return FormatCelsius(result)
}

func ConvertKToF(temp float64) string {
	result := (temp-273.15)*1.8 + 32
	return FormatCelsius(result)
}
