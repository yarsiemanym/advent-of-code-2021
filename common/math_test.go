package common

import "testing"

func Test_SumInt(t *testing.T) {
	sum := SumInt(1, 2, 3, 4)

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
