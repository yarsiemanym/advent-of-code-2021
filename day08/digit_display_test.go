package day08

import "testing"

func Test_digitDisplay_Zero(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(zero)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 0 {
		t.Errorf("Expected 0 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_One(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(one)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 1 {
		t.Errorf("Expected 1 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Two(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(two)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 2 {
		t.Errorf("Expected 2 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Three(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(three)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 3 {
		t.Errorf("Expected 3 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Four(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(four)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 4 {
		t.Errorf("Expected 4 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Five(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(five)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 5 {
		t.Errorf("Expected 5 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Six(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(six)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 6 {
		t.Errorf("Expected 6 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Seven(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(seven)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 7 {
		t.Errorf("Expected 7 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Eight(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(eight)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 8 {
		t.Errorf("Expected 8 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Nine(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals(nine)
	numericValue := digit.NumericValue()

	if numericValue == nil {
		t.Error("numericValue is nil.")
	} else if *numericValue != 9 {
		t.Errorf("Expected 9 but got %v.", *numericValue)
	}
}

func Test_digitDisplay_Nil(t *testing.T) {
	digit := NewDigitDisplay(0)
	digit.SetSignals("")
	numericValue := digit.NumericValue()

	if numericValue != nil {
		t.Error("numericValue is not nil.")
	}
}
