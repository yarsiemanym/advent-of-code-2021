package day21

import "testing"

func Test_Player_TakeTurn(t *testing.T) {
	player1 := NewPlayer(1, 4)
	player2 := NewPlayer(2, 8)
	dice := NewDeterministicDie(100)
	player1.TakeTurn(dice)
	player2.TakeTurn(dice)
	player1.TakeTurn(dice)
	player2.TakeTurn(dice)

	if player1.position != 4 {
		t.Errorf("Expected 4 but got %d.", player1.position)
	}

	if player1.score != 14 {
		t.Errorf("Expected 14 but got %d.", player1.position)
	}

	if player2.position != 6 {
		t.Errorf("Expected 6 but got %d.", player2.position)
	}

	if player2.score != 9 {
		t.Errorf("Expected 9 but got %d.", player2.position)
	}
}
