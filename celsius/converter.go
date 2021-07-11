package celsius

import "dojo-temp-converter/util"

type Celsius struct{}

func (c Celsius) ConvertToF(temp float64) string {
	result := 1.8*temp + 32
	return util.FormatFahrenheit(result)
}

func (c Celsius) ConvertToK(temp float64) string {
	result := temp + 273.15
	return util.FormatKelvin(result)
}

func (c Celsius) ConvertToC(temp float64) string {
	return util.FormatCelsius(temp)
}
