package common

import "fmt"

type Point struct {
	x int
	y int
	z int
}

func New2DPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
		z: 0,
	}
}

func New3DPoint(x int, y int, z int) *Point {
	return &Point{
		x: x,
		y: y,
		z: z,
	}
}

func (point *Point) X() int {
	return point.x
}

func (point *Point) Y() int {
	return point.y
}

func (point *Point) Z() int {
	return point.z
}

func (point *Point) Move(slope *Point) *Point {
	return &Point{
		x: point.X() + slope.X(),
		y: point.Y() + slope.Y(),
		z: point.Z() + slope.Z(),
	}
}

func (point *Point) ManhattanDistance(otherPoint *Point) int {
	return AbsInt(point.x-otherPoint.x) + AbsInt(point.y-otherPoint.y) + AbsInt(point.z-otherPoint.z)
}

func (point *Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
}
