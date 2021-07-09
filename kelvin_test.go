package dojotempconverter

import "testing"

func TestConvertKToC(t *testing.T) {
	Assert(ConvertKToC(250), "23.150", t)
}

func TestConvertKToF(t *testing.T) {
	Assert(ConvertKToF(255.37), "-0.004", t)
}
