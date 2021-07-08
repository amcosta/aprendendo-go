package dojotempconverter

import "testing"

func TestConvertToF(t *testing.T) {
	actual := ConvertToF(37)
	if actual != "98.60" {
		t.Errorf("The result %s mismatch with expected 98.6", actual)
	}
}

func TestConvertToK(t *testing.T) {
	actual := ConvertToK(37)
	if actual != "310.15" {
		t.Errorf("The result %s mismatch with expected 310.15", actual)
	}
}
