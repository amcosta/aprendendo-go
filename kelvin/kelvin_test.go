package kelvin

import "testing"

func assert(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("The result %s mismatch with expected %s", actual, expected)
	}
}

func TestConvertKToC(t *testing.T) {
	assert(ConvertKToC(250), "23.150", t)
}

func TestConvertKToF(t *testing.T) {
	assert(ConvertKToF(255.37), "-0.004", t)
}