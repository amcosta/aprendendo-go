package converter

import (
	"testing"
)

func celsiusAssert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertCToF(t *testing.T) {
	converter := new(Celsius)
	celsiusAssert(converter.ConvertToF(37), "98.600", t)
}

func TestConvertCToK(t *testing.T) {
	converter := new(Celsius)
	celsiusAssert(converter.ConvertToK(37), "310.150", t)
}
