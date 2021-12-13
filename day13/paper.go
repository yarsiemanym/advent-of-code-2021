package day13

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type Paper struct {
	plane *common.BoundedPlane
}

func NewPaper(height int, width int) *Paper {
	return &Paper{
		plane: common.NewBoundedPlane(height, width),
	}
}

func NewPaperFromPoints(points []*common.Point) *Paper {
	paper := &Paper{
		plane: common.NewBoundedPlaneFromPoints(points),
	}

	for _, point := range points {
		paper.DrawMark(point)
	}

	return paper
}

func (paper *Paper) Height() int {
	return paper.plane.Span().End().Y() + 1
}

func (paper *Paper) Width() int {
	return paper.plane.Span().End().X() + 1
}

func (paper *Paper) GetMarkAt(point *common.Point) rune {
	value := paper.plane.GetValueAt(point)

	if value == nil {
		return '.'
	} else {
		return value.(rune)
	}
}

func (paper *Paper) DrawMark(point *common.Point) {
	log.Tracef("Drawing mark on point (%d, %d).", point.X(), point.Y())
	paper.plane.SetValueAt(point, '#')
}

func (paper *Paper) Fold(crease *Crease) *Paper {
	log.Debugf("Folding paper along %c=%d.", crease.Axis, crease.Position)

	newHeight := paper.Height()
	newWidth := paper.Width()

	switch crease.Axis {
	case 'x':
		newWidth /= 2
	case 'y':
		newHeight /= 2
	default:
		log.Fatalf("'%c' is not a valid axis.", crease.Axis)
	}

	var foldedPaper *Paper = NewPaper(newHeight, newWidth)

	for _, point := range paper.GetMarkedPoints() {
		newX := point.X()
		newY := point.Y()

		if crease.Axis == 'x' && newX >= newWidth {
			newX = point.X() - (2 * common.MaxInt(0, point.X()-newWidth))
		}

		if crease.Axis == 'y' && newY >= newHeight {
			newY = point.Y() - (2 * common.MaxInt(0, point.Y()-newHeight))
		}

		foldedPoint := common.NewPoint(newX, newY)

		log.Tracef("Folding point (%d, %d) into (%d, %d). ", point.X(), point.Y(), foldedPoint.X(), foldedPoint.Y())
		mark := paper.GetMarkAt(point)
		if mark == '#' {
			foldedPaper.DrawMark(foldedPoint)
		}

	}

	return foldedPaper
}

func (paper *Paper) Render() string {
	output := ""

	for y := 0; y < paper.Height(); y++ {
		for x := 0; x < paper.Width(); x++ {
			point := common.NewPoint(x, y)
			output += string(paper.GetMarkAt(point))
		}
		output += "\n"
	}

	output = strings.Trim(output, "\n")

	return output
}

func (paper *Paper) GetMarkedPoints() []*common.Point {
	markedPoints := []*common.Point{}

	for _, point := range paper.plane.GetAllPoints() {
		mark := paper.GetMarkAt(point)
		if mark == '#' {
			markedPoints = append(markedPoints, point)
		}
	}

	return markedPoints
}
