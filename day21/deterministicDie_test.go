package day21

import "testing"

func Test_die_Roll(t *testing.T) {
	die := NewDeterministicDie(2)
	roll1 := die.Roll()

	if roll1 != 1 {
		t.Errorf("Expected 1 but got %d.", roll1)
	}

	roll2 := die.Roll()

	if roll2 != 2 {
		t.Errorf("Expected 2 but got %d.", roll2)
	}

	roll3 := die.Roll()

	if roll3 != 1 {
		t.Errorf("Expected 1 but got %d.", roll3)
	}
}
