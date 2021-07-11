package converter

import "dojo-temp-converter/util"

type Fahrenheit struct{}

func (f Fahrenheit) ConvertToC(temp float64) string {
	result := (temp - 32) / 1.8
	return util.FormatCelsius(result)
}

func (f Fahrenheit) ConvertToK(temp float64) string {
	result := (temp-32)/1.8 + 273.15
	return util.FormatKelvin(result)
}

func (f Fahrenheit) ConvertToF(temp float64) string {
	return util.FormatFahrenheit(temp)
}
