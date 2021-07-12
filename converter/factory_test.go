package converter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateCelsius(t *testing.T) {
	converter, _ := CreateConverter("C")
	actual := reflect.TypeOf(converter).String()
	expected := "*converter.Celsius"

	a, ok := converter.(Celsius)
	fmt.Println(a, ok)

	if actual != expected {
		t.Errorf("o tipo %s não é igual ao experado %s", actual, expected)
	}
}

func TestCreateFahrenheit(t *testing.T) {
	converter, _ := CreateConverter("F")
	actual := reflect.TypeOf(converter).String()
	expected := "*converter.Fahrenheit"

	if actual != expected {
		t.Errorf("o tipo %s não é igual ao experado %s", actual, expected)
	}
}

func TestCreateKelvin(t *testing.T) {
	converter, _ := CreateConverter("K")
	actual := reflect.TypeOf(converter).String()
	expected := "*converter.Kelvin"

	if actual != expected {
		t.Errorf("o tipo %s não é igual ao experado %s", actual, expected)
	}
}

func TestErrorOnInvalidType(t *testing.T) {
	converter, err := CreateConverter("A")
	if converter != nil && err == nil {
		t.Fail()
	}
}
