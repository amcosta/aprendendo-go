package fahrenheit

import "testing"

func assert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertFToC(t *testing.T) {
	assert(ConvertFToC(95), "35.000", t)
}

func TestConvertFToK(t *testing.T) {
	assert(ConvertFToK(59), "288.150", t)
}
