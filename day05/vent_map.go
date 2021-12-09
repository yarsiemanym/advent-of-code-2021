package day05

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type ventMap struct {
	plane common.BoundedPlane
}

func (ventMap *ventMap) Init(lines []*common.LineSegment) {
	ventMap.plane = *common.NewBoundedPlaneFromLines(lines)

	rows := ventMap.plane.Span().End().Y() + 1
	cols := ventMap.plane.Span().End().X() + 1

	ventMap.plane = *common.NewBoundedPlane(rows, cols)
}

func (ventMap *ventMap) ApplyLine(line *common.LineSegment) {
	log.Debugf("Applying line \"(%v,%v) -> (%v,%v)\" to the vent map.", line.Start().X(), line.Start().Y(), line.End().X(), line.End().Y())

	slope := line.Slope()
	log.Tracef("slope = (%v, %v)", slope.X(), slope.Y())

	if *line.Start() == *line.End() {
		value := ventMap.plane.GetValueAt(line.Start()).(int)
		ventMap.plane.SetValueAt(line.Start(), value+1)
	} else {
		for point := line.Start(); ; point = point.Move(slope) {

			log.Tracef("Incrementing overlaps at location (%v, %v).", point.X(), point.Y())
			value := ventMap.plane.GetValueAt(point)

			if value == nil {
				value = 0
			}

			ventMap.plane.SetValueAt(point, value.(int)+1)

			if *point == *line.End() {
				break
			}
		}
	}
}

func (ventMap *ventMap) CountOverlaps(withMinLines int) int {
	count := 0

	for _, point := range ventMap.plane.GetAllPoints() {
		value := ventMap.plane.GetValueAt(point)

		if value == nil {
			value = 0
		}

		if value.(int) >= withMinLines {
			log.Debugf("Location %v has %v overlapping lines. Incrementing count.", *point, value.(int))
			count++
		} else {
			log.Tracef("Location %v has %v overlapping lines. Skipping.", *point, value)
		}
	}

	return count
}

func (ventMap *ventMap) Render() string {
	message := "Vent Map\n"

	for row := ventMap.plane.Span().Start().Y(); row <= ventMap.plane.Span().End().Y(); row++ {
		for col := ventMap.plane.Span().Start().X(); col <= ventMap.plane.Span().End().X(); col++ {
			value := ventMap.plane.GetValueAt(common.NewPoint(col, row))
			if value == nil || value.(int) == 0 {
				message += "."
			} else {
				message += fmt.Sprintf("%v", value)
			}
		}
		message += "\n"
	}

	return message
}
