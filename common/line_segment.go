package common

type LineSegment struct {
	start *Point
	end   *Point
	slope *Point
}

func NewLineSegment(start *Point, end *Point) *LineSegment {
	return &LineSegment{
		start: start,
		end:   end,
	}
}

func (line *LineSegment) Start() *Point {
	return line.start
}

func (line *LineSegment) End() *Point {
	return line.end
}

func (line *LineSegment) Slope() *Point {
	if line.slope == nil {
		x := line.end.x - line.start.x
		y := line.end.y - line.start.y

		x, y = Reduce(x, y)

		line.slope = NewPoint(x, y)
	}

	return line.slope
}

func (line *LineSegment) IsVertical() bool {
	return line.start.x == line.end.x
}

func (line *LineSegment) IsHorizontal() bool {
	return line.start.y == line.end.y
}
