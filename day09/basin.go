package day09

import (
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type basin struct {
	points []*common.Point
}

func NewBasin() *basin {
	return &basin{}
}

func (basin *basin) Add(points ...*common.Point) {
	log.Tracef("Attempting to add %v points to basin.", len(points))
	for _, point := range points {
		if !basin.Contains(point) {
			log.Tracef("Adding point %v to basin.", *point)
			basin.points = append(basin.points, point)
		} else {
			log.Tracef("Point %v already exists in this basin. Skipping.", *point)
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
