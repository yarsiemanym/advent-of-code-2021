package day06

import "testing"

func Test_lanternfish_Tick_NoBaby(t *testing.T) {
	fish := &lanternfish{
		timer: 2,
	}

	baby := fish.Tick()

	if fish.timer != 1 {
		t.Errorf("Expected 1 but got %v.", fish.timer)
	}

	if baby != nil {
		t.Error("Baby is not nil.")
	}
}

func Test_lanternfish_Tick_Baby(t *testing.T) {
	parent := &lanternfish{
		timer: 0,
	}

	baby := parent.Tick()

	if parent.timer != 6 {
		t.Errorf("Expected 6 but got %v.", parent.timer)
	}

	if baby == nil {
		t.Error("Baby is nil.")
	} else if baby.timer != 8 {
		t.Errorf("Expected 8 but got %v.", baby.timer)
	}
}
