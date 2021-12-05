package day05

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type ventMap struct {
	plane     common.BoundedPlane
	Locations [][]int
}

func (ventMap *ventMap) Init(lines []*common.LineSegment) {
	ventMap.plane = *common.NewBoundedPlaneFromLines(lines)

	rows := ventMap.plane.Span().End().Y() + 1
	cols := ventMap.plane.Span().End().X() + 1

	ventMap.Locations = make([][]int, rows)

	for i := 0; i < rows; i++ {
		ventMap.Locations[i] = make([]int, cols)
	}
}

func (ventMap *ventMap) ApplyLine(line *common.LineSegment) {
	log.Debugf("Applying line \"(%v,%v) -> (%v,%v)\" to the vent map.", line.Start().X(), line.Start().Y(), line.End().X(), line.End().Y())

	slope := line.Slope()
	log.Tracef("slope = (%v, %v)", slope.X(), slope.Y())

	if *line.Start() == *line.End() {
		ventMap.Locations[line.Start().Y()][line.Start().X()] += 1
	} else {
		for point := line.Start(); ; point = point.Move(slope) {

			log.Tracef("Incrementing overlaps at location (%v, %v).", point.X(), point.Y())
			ventMap.Locations[point.Y()][point.X()] += 1

			if *point == *line.End() {
				break
			}
		}
	}
}

func (ventMap *ventMap) CountOverlaps(withMinLines int) int {
	count := 0

	for y, row := range ventMap.Locations {
		for x, location := range row {
			if location >= withMinLines {
				log.Debugf("Location (%v, %v) has %v overlapping lines. Incrementing count.", x, y, location)
				count++
			} else {
				log.Tracef("Location (%v, %v) has %v overlapping lines. Skipping.", x, y, location)
			}
		}
	}

	return count
}

func (ventMap *ventMap) Print() {
	if log.GetLevel() == log.TraceLevel {
		message := "Vent Map\n"

		for _, row := range ventMap.Locations {
			for _, location := range row {
				if location == 0 {
					message += "."
				} else {
					message += fmt.Sprintf("%v", location)
				}
			}
			message += "\n"
		}

		log.Trace(message)
	}
}
