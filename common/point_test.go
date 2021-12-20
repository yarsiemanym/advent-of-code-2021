package common

import "testing"

func Test_Point_Move(t *testing.T) {
	start := &Point{
		x: 1,
		y: 2,
		z: 3,
	}

	slope := &Point{
		x: 3,
		y: 4,
		z: -1,
	}

	end := start.Move(slope)

	if end.x != 4 {
		t.Errorf("Expected 4 but got %v.", end.x)
	}

	if end.y != 6 {
		t.Errorf("Expected 6 but got %v.", end.y)
	}

	if end.z != 2 {
		t.Errorf("Expected 2 but got %v.", end.z)
	}
}

func Test_Point_ManhattanDistance(t *testing.T) {
	point1 := New3DPoint(1, 2, -3)
	point2 := New3DPoint(-4, 0, 2)
	distance := point1.ManhattanDistance(point2)

	if distance != 12 {
		t.Errorf("Expected 12 but got %d.", distance)
	}
}
