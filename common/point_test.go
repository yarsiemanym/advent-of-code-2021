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

func Test_Point_ManhattanDistance(t *testing.T) {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(-4, 0)
	distance := point1.ManhattanDistance(point2)

	if distance != 7 {
		t.Errorf("Expected 7 but got %d.", distance)
	}
}
