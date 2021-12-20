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
}

func (ventMap *ventMap) GetOverlapsAt(point *common.Point) int {
	value := ventMap.plane.GetValueAt(point)

	if value == nil {
		value = 0
	}

	return value.(int)
}

func (ventMap *ventMap) ApplyLine(line *common.LineSegment) {
	log.Debugf("Applying line \"%s -> %s\" to the vent map.", line.Start(), line.End())

	slope := line.Slope()
	log.Tracef("slope = %s", slope)

	if *line.Start() == *line.End() {
		value := ventMap.GetOverlapsAt(line.Start())
		ventMap.plane.SetValueAt(line.Start(), value+1)
	} else {
		for point := line.Start(); ; point = point.Move(slope) {

			log.Tracef("Incrementing overlaps at location %s.", point)
			value := ventMap.GetOverlapsAt(point)
			ventMap.plane.SetValueAt(point, value+1)

			if *point == *line.End() {
				break
			}
		}
	}
}

func (ventMap *ventMap) CountOverlaps(withMinLines int) int {
	count := 0

	for _, point := range ventMap.plane.GetAllPoints() {
		value := ventMap.GetOverlapsAt(point)

		if value >= withMinLines {
			log.Debugf("Location %s has %d overlapping lines. Incrementing count.", point, value)
			count++
		} else {
			log.Tracef("Location %s has %d overlapping lines. Skipping.", point, value)
		}
	}

	return count
}

func (ventMap *ventMap) Render() string {
	message := "Vent Map\n"

	for row := ventMap.plane.Span().Start().Y(); row <= ventMap.plane.Span().End().Y(); row++ {
		for col := ventMap.plane.Span().Start().X(); col <= ventMap.plane.Span().End().X(); col++ {
			value := ventMap.GetOverlapsAt(common.New2DPoint(col, row))
			if value == 0 {
				message += "."
			} else {
				message += fmt.Sprintf("%v", value)
			}
		}
		message += "\n"
	}

	return message
}
