package dojotempconverter

import "testing"

func TestConvertFToC(t *testing.T) {
	Assert(ConvertFToC(95), "35.000", t)
}

func TestConvertFToK(t *testing.T) {
	Assert(ConvertFToK(59), "288.150", t)
}
