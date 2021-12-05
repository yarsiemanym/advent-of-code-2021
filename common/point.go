package common

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (point *Point) X() int {
	return point.x
}

func (point *Point) Y() int {
	return point.y
}

func (point *Point) Move(slope *Point) *Point {
	return &Point{
		x: point.X() + slope.X(),
		y: point.Y() + slope.Y(),
	}
}
