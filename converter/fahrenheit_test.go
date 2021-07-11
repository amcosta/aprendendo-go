package converter

import (
	"testing"
)

func fAssert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertFToC(t *testing.T) {
	converter := new(Fahrenheit)
	fAssert(converter.ConvertToC(95), "35.000", t)
}

func TestConvertFToK(t *testing.T) {
	converter := new(Fahrenheit)
	fAssert(converter.ConvertToK(59), "288.150", t)
}
