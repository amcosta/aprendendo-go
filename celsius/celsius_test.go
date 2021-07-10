package celsius

import "testing"

func assert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertCToF(t *testing.T) {
	assert(ConvertToF(37), "98.600", t)
}

func TestConvertCToK(t *testing.T) {
	assert(ConvertToK(37), "310.150", t)
}
