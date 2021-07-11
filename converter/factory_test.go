package converter

import (
	"reflect"
	"testing"
)

func TestCreateCelsius(t *testing.T) {
	actual, _ := CreateConverter("C")

	a := reflect.Struct
	if a != "converter.Celsius" {
		t.Fail()
	}
}
