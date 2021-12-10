package day09

import "github.com/yarsiemanym/advent-of-code-2021/common"

type basin struct {
	points []*common.Point
}

func NewBasin() *basin {
	return &basin{}
}

func (basin *basin) Add(points ...*common.Point) {
	for _, point := range points {
		if !basin.Contains(point) {
			basin.points = append(basin.points, point)
		}
	}
}

func (basin *basin) Contains(point *common.Point) bool {
	for _, member := range basin.points {
		if *member == *point {
			return true
		}
	}

	return false
}

func (basin *basin) Size() int {
	return len(basin.points)
}
