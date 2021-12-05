package common

import "testing"

func Test_SumIntVariadic(t *testing.T) {
	sum := SumIntVariadic(1, 2, 3, 4)

	if sum != 10 {
		t.Errorf("Expected 10 but got %v.", sum)
	}
}

func Test_SumIntSlice(t *testing.T) {
	values := []int{1, 2, 3, 4}

	sum := SumIntSlice(values)

	if sum != 10 {
		t.Errorf("Expected 10 but got %v.", sum)
	}
}

func Test_MaxIntVariadic(t *testing.T) {
	max := MaxIntVariadic(3, 1, 4, 2)

	if max != 4 {
		t.Errorf("Expected 4 but got %v.", max)
	}
}

func Test_MaxIntSlice(t *testing.T) {
	values := []int{3, 1, 4, 2}

	max := MaxIntSlice(values)

	if max != 4 {
		t.Errorf("Expected 4 but got %v.", max)
	}
}

func Test_MinIntVariadic(t *testing.T) {
	min := MinIntVariadic(3, 1, 4, 2)

	if min != 1 {
		t.Errorf("Expected 1 but got %v.", min)
	}
}

func Test_MinIntSlice(t *testing.T) {
	values := []int{3, 1, 4, 2}

	min := MinIntSlice(values)

	if min != 1 {
		t.Errorf("Expected 1 but got %v.", min)
	}
}

func Test_GreatestCommonDenominator_NotOne(t *testing.T) {
	gdc := GreatestCommonDenominator(15, 6)

	if gdc != 3 {
		t.Errorf("Expected 3 but got %v.", gdc)
	}
}

func Test_GreatestCommonDenominator_One(t *testing.T) {
	gdc := GreatestCommonDenominator(19, 7)

	if gdc != 1 {
		t.Errorf("Expected 1 but got %v.", gdc)
	}
}

func Test_GreatestCommonDenominator_Negative(t *testing.T) {
	gdc := GreatestCommonDenominator(-15, 6)

	if gdc != 3 {
		t.Errorf("Expected 3 but got %v.", gdc)
	}
}

func Test_GreatestCommonDenominator_Zero(t *testing.T) {
	gdc := GreatestCommonDenominator(0, 5)

	if gdc != 1 {
		t.Errorf("Expected 1 but got %v.", gdc)
	}
}

func Test_Reduce(t *testing.T) {
	numerator, denominator := Reduce(-15, 6)

	if numerator != -5 {
		t.Errorf("Expected -5 but got %v.", numerator)
	}

	if denominator != 2 {
		t.Errorf("Expected 2 but got %v.", denominator)
	}
}
