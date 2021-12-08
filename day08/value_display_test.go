package day08

import "testing"

func Test_valueDisplay_1234(t *testing.T) {
	value := NewValueDisplay()
	value.SetSignals(one, two, three, four, nil)

	numerivValue := value.NumericValue()

	if numerivValue == nil {
		t.Error("numericValue is nil.")
	} else if *numerivValue != 1234 {
		t.Errorf("Expected 1234 but got %v.", *numerivValue)
	}
}

func Test_valueDisplay_Mapped(t *testing.T) {
	signalPatterns := []string{
		"acedgfb",
		"cdfbe",
		"gcdfa",
		"fbcad",
		"dab",
		"cefabd",
		"cdfgeb",
		"eafb",
		"cagedb",
		"ab",
	}

	signalMapper := NewSignalMapper(signalPatterns)

	display := NewValueDisplay()
	display.SetSignals("cdfeb", "fcadb", "cdfeb", "cdbaf", signalMapper)
	value := display.NumericValue()

	if value == nil {
		t.Error("value is nil.")
	} else if *value != 5353 {
		t.Errorf("Expected 5353 but got %v.", *value)
	}
}
