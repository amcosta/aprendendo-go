package fahrenheit

import (
	"dojo-temp-converter/util"
	"testing"
)

var converter util.Converter

func init() {
	converter = new(Fahrenheit)
}

func assert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertFToC(t *testing.T) {
	assert(converter.ConvertToC(95), "35.000", t)
}

func TestConvertFToK(t *testing.T) {
	assert(converter.ConvertToK(59), "288.150", t)
}
