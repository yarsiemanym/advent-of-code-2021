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
		minX = MinInt(minX, point.x)
		maxX = MaxInt(maxX, point.x)
		minY = MinInt(minY, point.y)
		maxY = MaxInt(maxY, point.y)
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
		minX = MinInt(minX, line.start.x, line.end.x)
		maxX = MaxInt(maxX, line.start.x, line.end.x)
		minY = MinInt(minY, line.start.y, line.end.y)
		maxY = MaxInt(maxY, line.start.y, line.end.y)
	}

	return &BoundedPlane{
		span: NewLineSegment(NewPoint(minX, minY), NewPoint(maxX, maxY)),
	}
}

func (plane *BoundedPlane) Span() *LineSegment {
	return plane.span
}
