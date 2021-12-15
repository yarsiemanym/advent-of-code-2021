package day15

import (
	"fmt"
	"strings"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type Path struct {
	points []*common.Point
}

func NewPath() *Path {
	return &Path{}
}

func (path *Path) Clone() *Path {
	clone := &Path{
		points: make([]*common.Point, len(path.points)),
	}
	copy(clone.points, path.points)
	return clone
}

func (path *Path) Append(point *common.Point) {
	path.points = append(path.points, point)
}

func (path *Path) Prepend(point *common.Point) {
	path.points = append([]*common.Point{point}, path.points...)
}

func (path *Path) Contains(point *common.Point) bool {
	exists := false

	for _, pathPoint := range path.points {
		if *point == *pathPoint {
			exists = true
			break
		}
	}

	return exists
}

func (path *Path) Overlaps(otherPath *Path) bool {
	for _, point := range otherPath.points {
		if path.Contains(point) {
			return true
		}
	}

	return false
}

func (path *Path) Start() *common.Point {
	if len(path.points) > 0 {
		return path.points[0]
	}

	return nil
}

func (path *Path) Points() []*common.Point {
	return path.points
}

func (path *Path) End() *common.Point {
	if len(path.points) > 0 {
		return path.points[len(path.points)-1]
	}

	return nil
}

func (path *Path) String() string {
	output := "["

	for _, point := range path.points {
		output += fmt.Sprintf("%s => ", point)
	}

	output = strings.Trim(output, " =>") + "]"
	return output
}
