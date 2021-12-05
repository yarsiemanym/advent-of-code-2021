package common

import "testing"

func Test_Point_Move(t *testing.T) {
	start := &Point{
		x: 1,
		y: 2,
	}

	slope := &Point{
		x: 3,
		y: 4,
	}

	end := start.Move(slope)

	if end.x != 4 {
		t.Errorf("Expected 4 but got %v.", end.x)
	}

	if end.y != 6 {
		t.Errorf("Expected 6 but got %v.", end.y)
	}
}
