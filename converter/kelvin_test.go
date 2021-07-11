package converter

import (
	"testing"
)

func kAssert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertKToC(t *testing.T) {
	converter := new(Kelvin)
	kAssert(converter.ConvertToC(250), "23.150", t)
}

func TestConvertKToF(t *testing.T) {
	converter := new(Kelvin)
	kAssert(converter.ConvertToF(255.37), "-0.004", t)
}

func TestConvertKToK(t *testing.T) {
	converter := new(Kelvin)
	kAssert(converter.ConvertToK(230.37), "230.370", t)
}
