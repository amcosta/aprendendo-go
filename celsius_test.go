package dojotempconverter

import "testing"

func TestConvertCToF(t *testing.T) {
	Assert(ConvertCToF(37), "98.600", t)
}

func TestConvertCToK(t *testing.T) {
	Assert(ConvertCToK(37), "310.150", t)
}
