package common

type BoundedPlane struct {
	span *LineSegment
}

func NewBoundedPlaneFromPoints(points []*Point) *BoundedPlane {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for _, point := range points {
		minX = MinIntVariadic(minX, point.x)
		maxX = MaxIntVariadic(maxX, point.x)
		minY = MinIntVariadic(minY, point.y)
		maxY = MaxIntVariadic(maxY, point.y)
	}

	return &BoundedPlane{
		span: NewLineSegment(NewPoint(minX, minY), NewPoint(maxX, maxY)),
	}
}

func NewBoundedPlaneFromLines(lines []*LineSegment) *BoundedPlane {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for _, line := range lines {
		minX = MinIntVariadic(minX, line.start.x, line.end.x)
		maxX = MaxIntVariadic(maxX, line.start.x, line.end.x)
		minY = MinIntVariadic(minY, line.start.y, line.end.y)
		maxY = MaxIntVariadic(maxY, line.start.y, line.end.y)
	}

	return &BoundedPlane{
		span: NewLineSegment(NewPoint(minX, minY), NewPoint(maxX, maxY)),
	}
}

func (plane *BoundedPlane) Span() *LineSegment {
	return plane.span
}
