package celsius

import (
	"dojo-temp-converter/util"
	"testing"
)

var converter util.Converter

func init() {
	converter = new(Celsius)
}

func assert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertCToF(t *testing.T) {
	assert(converter.ConvertToF(37), "98.600", t)
}

func TestConvertCToK(t *testing.T) {
	assert(converter.ConvertToK(37), "310.150", t)
}
