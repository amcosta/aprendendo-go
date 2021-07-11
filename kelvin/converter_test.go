package kelvin

import (
	"dojo-temp-converter/util"
	"testing"
)

var converter util.Converter

func init() {
	converter = new(Kelvin)
}

func assert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertKToC(t *testing.T) {
	assert(converter.ConvertToC(250), "23.150", t)
}

func TestConvertKToF(t *testing.T) {
	assert(converter.ConvertToF(255.37), "-0.004", t)
}
