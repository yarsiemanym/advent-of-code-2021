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
	if path == nil {
		return nil
	}

	clone := &Path{
		points: make([]*common.Point, len(path.points)),
	}
	copy(clone.points, path.points)
	return clone
}

func (path *Path) Append(point *common.Point) {
	if path == nil {
		return
	}

	path.points = append(path.points, point)
}

func (path *Path) Prepend(point *common.Point) {
	if path == nil {
		return
	}

	path.points = append([]*common.Point{point}, path.points...)
}

func (path *Path) Contains(point *common.Point) bool {
	if path == nil {
		return false
	}

	exists := false

	for _, pathPoint := range path.points {
		if *point == *pathPoint {
			exists = true
			break
		}
	}

	return exists
}

func (path *Path) Length() int {
	if path == nil {
		return 0
	}

	return len(path.points)
}

func (path *Path) Start() *common.Point {
	if path == nil || len(path.points) == 0 {
		return nil
	}

	return path.points[0]
}

func (path *Path) Points() []*common.Point {
	if path == nil {
		return nil
	}

	return path.points
}

func (path *Path) End() *common.Point {
	if path == nil || len(path.points) == 0 {
		return nil
	}

	return path.points[len(path.points)-1]
}

func (path *Path) String() string {
	if path == nil {
		return "[]"
	}

	output := "["

	for _, point := range path.points {
		output += fmt.Sprintf("%s => ", point)
	}

	output = strings.Trim(output, " =>") + "]"
	return output
}
