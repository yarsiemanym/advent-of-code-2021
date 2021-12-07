package day06

import "testing"

func Test_lanternfish_DecedantsAfterOneDay_NoDecendants(t *testing.T) {
	parent := &lanternfish{
		timer: 2,
	}

	decendants := parent.DecendantsAfter(1)

	if decendants != 0 {
		t.Errorf("Expected 0 but got %v.", decendants)
	}
}

func Test_lanternfish_DecendantsAfterOneDay_OneDecendants(t *testing.T) {
	parent := &lanternfish{
		timer: 0,
	}

	decendants := parent.DecendantsAfter(1)

	if decendants != 1 {
		t.Errorf("Expected 1 but got %v.", decendants)
	}
}

func Test_lanternfish_DecendantsAfterEightDays_TwoDecendants(t *testing.T) {
	parent := &lanternfish{
		timer: 0,
	}

	decendants := parent.DecendantsAfter(8)

	if decendants != 2 {
		t.Errorf("Expected 2 but got %v.", decendants)
	}
}

func Test_lanternfish_DecendantsAfterSixteenDays_FourDecendants(t *testing.T) {
	parent := &lanternfish{
		timer: 1,
	}

	decendants := parent.DecendantsAfter(16)

	if decendants != 4 {
		t.Errorf("Expected 4 but got %v.", decendants)
	}
}

