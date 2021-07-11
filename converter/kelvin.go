package converter

import "dojo-temp-converter/util"

type Kelvin struct{}

func (k Kelvin) ConvertToC(temp float64) string {
	result := temp*-1 + 273.15
	return util.FormatCelsius(result)
}

func (k Kelvin) ConvertToF(temp float64) string {
	result := (temp-273.15)*1.8 + 32
	return util.FormatCelsius(result)
}

func (k Kelvin) ConvertToK(temp float64) string {
	return util.FormatKelvin(temp)
}
