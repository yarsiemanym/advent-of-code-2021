package day06

import "testing"

func Test_lanternfish_LineageAfterOneDay_NoAncestors(t *testing.T) {
	fish := &lanternfish{
		timer: 2,
	}

	lineage := fish.AncestorsAfter(1)

	if lineage != 1 {
		t.Errorf("Expected 1 but got %v.", lineage)
	}
}

func Test_lanternfish_LineageAfterOneDay_TwoFish(t *testing.T) {
	parent := &lanternfish{
		timer: 0,
	}

	lineage := parent.AncestorsAfter(1)

	if lineage != 2 {
		t.Errorf("Expected 2 but got %v.", lineage)
	}
}

func Test_lanternfish_LineageAfterEightDays_ThreeFish(t *testing.T) {
	parent := &lanternfish{
		timer: 0,
	}

	lineage := parent.AncestorsAfter(8)

	if lineage != 3 {
		t.Errorf("Expected 3 but got %v.", lineage)
	}
}

func Test_lanternfish_LineageAfterSixteenDays_FiveFish(t *testing.T) {
	parent := &lanternfish{
		timer: 1,
	}

	lineage := parent.AncestorsAfter(16)

	if lineage != 5 {
		t.Errorf("Expected 5 but got %v.", lineage)
	}
}

func Test_lanternfish_LineageAfterSeventeenDays_FiveFish(t *testing.T) {
	parent := &lanternfish{
		timer: 1,
	}

	lineage := parent.AncestorsAfter(17)

	if lineage != 5 {
		t.Errorf("Expected 5 but got %v.", lineage)
	}
}
