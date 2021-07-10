package fahrenheit

import "testing"

func assert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertFToC(t *testing.T) {
	assert(ConvertToC(95), "35.000", t)
}

func TestConvertFToK(t *testing.T) {
	assert(ConvertToK(59), "288.150", t)
}
