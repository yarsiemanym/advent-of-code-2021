package common

type BoundedPlane struct {
	span      *LineSegment
	locations [][]interface{}
}

func NewBoundedPlane(height int, width int) *BoundedPlane {
	return &BoundedPlane{
		span:      NewLineSegment(NewPoint(0, 0), NewPoint(width-1, height-1)),
		locations: initializeLocations(height, width),
	}
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
		span:      NewLineSegment(NewPoint(minX, minY), NewPoint(maxX, maxY)),
		locations: initializeLocations(maxY-minY, maxX-minX),
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
		span:      NewLineSegment(NewPoint(minX, minY), NewPoint(maxX, maxY)),
		locations: initializeLocations(maxY-minY, maxX-minX),
	}
}

func initializeLocations(height int, width int) [][]interface{} {
	locations := make([][]interface{}, height)

	for row := range locations {

		locations[row] = make([]interface{}, width)
	}

	return locations
}

func (plane *BoundedPlane) Span() *LineSegment {
	return plane.span
}

func (plane *BoundedPlane) GetValueAt(point *Point) interface{} {
	row := point.y - plane.span.start.y
	col := point.x - plane.span.start.x
	return plane.locations[row][col]
}

func (plane *BoundedPlane) SetValueAt(point *Point, value interface{}) {
	row := point.y - plane.span.start.y
	col := point.x - plane.span.start.x
	plane.locations[row][col] = value
}

func (plane *BoundedPlane) GetPointsOrthoganallyAdjacentTo(point *Point) []*Point {
	var adjacentPoints []*Point

	if point.x+1 <= plane.span.end.x {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x+1, point.y))
	}

	if point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x, point.y+1))
	}

	if point.x-1 >= plane.span.start.x {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x-1, point.y))
	}

	if point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x, point.y-1))
	}

	return adjacentPoints
}

func (plane *BoundedPlane) GetPointsAdjacentTo(point *Point) []*Point {
	var adjacentPoints []*Point

	if point.x+1 <= plane.span.end.x {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x+1, point.y))
	}

	if point.x+1 <= plane.span.end.x && point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x+1, point.y+1))
	}

	if point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x, point.y+1))
	}

	if point.x-1 >= plane.span.start.x && point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x-1, point.y+1))
	}

	if point.x-1 >= plane.span.start.x {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x-1, point.y))
	}

	if point.x-1 >= plane.span.start.x && point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x-1, point.y-1))
	}

	if point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x, point.y-1))
	}

	if point.x+1 <= plane.span.end.x && point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, NewPoint(point.x+1, point.y-1))
	}

	return adjacentPoints
}

func (plane *BoundedPlane) GetAllPoints() []*Point {
	var points []*Point

	for y := plane.span.start.y; y <= plane.span.end.y; y++ {
		for x := plane.span.start.x; x <= plane.span.end.x; x++ {
			points = append(points, NewPoint(x, y))
		}
	}

	return points
}
