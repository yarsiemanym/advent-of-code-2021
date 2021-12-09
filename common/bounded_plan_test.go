package common

import "testing"

func Test_NewBoundedPlaneFromPoints(t *testing.T) {
	points := []*Point{
		NewPoint(3, 3),
		NewPoint(-3, 3),
		NewPoint(-3, -3),
		NewPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	if plane == nil {
		t.Error("plane is nil")
	} else if plane.span == nil {
		t.Error("plane.span is nil")
	} else if plane.span.start == nil {
		t.Error("plane.span.start is nil")
	} else if plane.span.start.x != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.x)
	} else if plane.span.start.y != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.y)
	} else if plane.span.end.x != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.x)
	} else if plane.span.end.y != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.y)
	}

	if len(plane.locations) != 6 {
		t.Errorf("Expected 6 but got %v.", len(plane.locations))
	}

	for _, row := range plane.locations {
		if len(row) != 6 {
			t.Errorf("Expected 6 but got %v.", len(row))
		}
	}
}

func Test_NewBoundedPlaneFromLines(t *testing.T) {
	lines := []*LineSegment{
		NewLineSegment(NewPoint(3, 1), NewPoint(1, -3)),
		NewLineSegment(NewPoint(-3, -1), NewPoint(-1, 3)),
	}

	plane := NewBoundedPlaneFromLines(lines)

	if plane == nil {
		t.Error("plane is nil")
	} else if plane.span == nil {
		t.Error("plane.span is nil")
	} else if plane.span.start == nil {
		t.Error("plane.span.start is nil")
	} else if plane.span.start.x != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.x)
	} else if plane.span.start.y != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.y)
	} else if plane.span.end.x != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.x)
	} else if plane.span.end.y != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.y)
	}

	if len(plane.locations) != 6 {
		t.Errorf("Expected 6 but got %v.", len(plane.locations))
	}

	for _, row := range plane.locations {
		if len(row) != 6 {
			t.Errorf("Expected 6 but got %v.", len(row))
		}
	}
}

func Test_boundedPlane_GetValueAt(t *testing.T) {
	points := []*Point{
		NewPoint(3, 3),
		NewPoint(-3, 3),
		NewPoint(-3, -3),
		NewPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)
	plane.locations[1][1] = true

	value1 := plane.GetValueAt(NewPoint(-2, -2)).(bool)
	value2 := plane.GetValueAt(NewPoint(1, 1))

	if !value1 {
		t.Errorf("Expected true but got %v.", value1)
	}

	if value2 != nil {
		t.Error("value2 is not nil.")
	}
}

func Test_boundedPlane_SetValueAt(t *testing.T) {
	points := []*Point{
		NewPoint(3, 3),
		NewPoint(-3, 3),
		NewPoint(-3, -3),
		NewPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)
	plane.SetValueAt(NewPoint(-2, -2), true)

	value1 := plane.locations[1][1].(bool)
	value2 := plane.GetValueAt(NewPoint(1, 1))

	if !value1 {
		t.Errorf("Expected true but got %v.", value1)
	}

	if value2 != nil {
		t.Error("value2 is not nil.")
	}
}

func Test_boundedPlane_GetPointsAdjacentTo_Interior(t *testing.T) {
	points := []*Point{
		NewPoint(3, 3),
		NewPoint(-3, 3),
		NewPoint(-3, -3),
		NewPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacenPoints := plane.GetPointsAdjacentTo(NewPoint(1, 1))

	if len(adjacenPoints) != 4 {
		t.Errorf("Expected 4 but got %v.", len(adjacenPoints))
	}

	if adjacenPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacenPoints[0].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacenPoints[0].x)
	} else if adjacenPoints[0].y != 1 {
		t.Errorf("Expected 1 but got %v.", adjacenPoints[0].y)
	}

	if adjacenPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacenPoints[1].x != 1 {
		t.Errorf("Expected 1 but got %v.", adjacenPoints[1].x)
	} else if adjacenPoints[1].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacenPoints[1].y)
	}

	if adjacenPoints[2] == nil {
		t.Error("adjacenPoints[2] is nil.")
	} else if adjacenPoints[2].x != 0 {
		t.Errorf("Expected 0 but got %v.", adjacenPoints[2].x)
	} else if adjacenPoints[2].y != 1 {
		t.Errorf("Expected 1 but got %v.", adjacenPoints[2].y)
	}

	if adjacenPoints[3] == nil {
		t.Error("adjacenPoints[3] is nil.")
	} else if adjacenPoints[3].x != 1 {
		t.Errorf("Expected 1 but got %v.", adjacenPoints[3].x)
	} else if adjacenPoints[3].y != 0 {
		t.Errorf("Expected 0 but got %v.", adjacenPoints[3].y)
	}
}

func Test_boundedPlane_GetPointsAdjacentTo_Corner1(t *testing.T) {
	points := []*Point{
		NewPoint(3, 3),
		NewPoint(-3, 3),
		NewPoint(-3, -3),
		NewPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacenPoints := plane.GetPointsAdjacentTo(NewPoint(-3, -3))

	if len(adjacenPoints) != 2 {
		t.Errorf("Expected 2 but got %v.", len(adjacenPoints))
	}

	if adjacenPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacenPoints[0].x != -2 {
		t.Errorf("Expected -2 but got %v.", adjacenPoints[0].x)
	} else if adjacenPoints[0].y != -3 {
		t.Errorf("Expected -2 but got %v.", adjacenPoints[0].y)
	}

	if adjacenPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacenPoints[1].x != -3 {
		t.Errorf("Expected -3 but got %v.", adjacenPoints[1].x)
	} else if adjacenPoints[1].y != -2 {
		t.Errorf("Expected -2 but got %v.", adjacenPoints[1].y)
	}
}

func Test_boundedPlane_GetPointsAdjacentTo_Corner2(t *testing.T) {
	points := []*Point{
		NewPoint(3, 3),
		NewPoint(-3, 3),
		NewPoint(-3, -3),
		NewPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacenPoints := plane.GetPointsAdjacentTo(NewPoint(3, 3))

	if len(adjacenPoints) != 2 {
		t.Errorf("Expected 2 but got %v.", len(adjacenPoints))
	}

	if adjacenPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacenPoints[0].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacenPoints[0].x)
	} else if adjacenPoints[0].y != 3 {
		t.Errorf("Expected 2 but got %v.", adjacenPoints[0].y)
	}

	if adjacenPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacenPoints[1].x != 3 {
		t.Errorf("Expected 3 but got %v.", adjacenPoints[1].x)
	} else if adjacenPoints[1].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacenPoints[1].y)
	}
}
