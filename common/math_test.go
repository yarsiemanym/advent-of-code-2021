package common

import "testing"

func Test_SumInt(t *testing.T) {
	sum := SumInt(1, 2, 3, 4)

	if sum != 10 {
		t.Errorf("Expected 10 but got %v.", sum)
	}
}

func Test_MaxInt(t *testing.T) {
	max := MaxInt(3, 1, 4, 2)

	if max != 4 {
		t.Errorf("Expected 4 but got %v.", max)
	}
}

func Test_MinInt(t *testing.T) {
	min := MinInt(3, 1, 4, 2)

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

func Test_PowInt(t *testing.T) {
	result := PowInt(2, 3)

	if result != 8 {
		t.Errorf("Expected 8 but got %v.", result)
	}
}

func Test_MedianInt_Empty(t *testing.T) {
	assertPanic(t, func() { MedianInt([]int{}...) })
}

func Test_MedianInt_Odd(t *testing.T) {
	values := []int{1, 2, 3}
	median := MedianInt(values...)

	if median != 2 {
		t.Errorf("Expected 2 but got %v.", median)
	}
}

func Test_MedianInt_Even(t *testing.T) {
	values := []int{1, 2, 3, 4}
	median := MedianInt(values...)

	if median != 3 {
		t.Errorf("Expected 3 but got %v.", median)
	}
}
