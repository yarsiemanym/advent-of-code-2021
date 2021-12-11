package common

import "testing"

func Test_LineSegment_Slope_Positive(t *testing.T) {
	line := LineSegment{
		start: &Point{
			x: 1,
			y: 2,
		},
		end: &Point{
			x: 5,
			y: 4,
		},
	}

	slope := line.Slope()

	if slope.x != 2 {
		t.Errorf("Expected 2 but got %v.", slope.x)
	}

	if slope.y != 1 {
		t.Errorf("Expected 1 but got %v.", slope.y)
	}
}

func Test_LineSegment_Slope_Negative(t *testing.T) {
	line := LineSegment{
		start: &Point{
			x: 5,
			y: 4,
		},
		end: &Point{
			x: 1,
			y: 2,
		},
	}

	slope := line.Slope()

	if slope.x != -2 {
		t.Errorf("Expected -2 but got %v.", slope.x)
	}

	if slope.y != -1 {
		t.Errorf("Expected -1 but got %v.", slope.y)
	}
}

func Test_LineSegment_Slope_Horizontal(t *testing.T) {
	line := LineSegment{
		start: &Point{
			x: 1,
			y: 4,
		},
		end: &Point{
			x: 3,
			y: 4,
		},
	}

	slope := line.Slope()

	if slope.x != 1 {
		t.Errorf("Expected 1 but got %v.", slope.x)
	}

	if slope.y != 0 {
		t.Errorf("Expected 0 but got %v.", slope.y)
	}
}

func Test_LineSegment_Slope_Vertical(t *testing.T) {
	line := LineSegment{
		start: &Point{
			x: 1,
			y: 2,
		},
		end: &Point{
			x: 1,
			y: 4,
		},
	}

	slope := line.Slope()

	if slope.x != 0 {
		t.Errorf("Expected 0 but got %v.", slope.x)
	}

	if slope.y != 1 {
		t.Errorf("Expected 1 but got %v.", slope.y)
	}
}
