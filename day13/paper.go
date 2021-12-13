package day13

import (
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

	var foldedPaper *Paper = nil

	switch crease.Axis {
	case 'x':
		foldedPaper = paper.foldLeft(crease.Position)
	case 'y':
		foldedPaper = paper.foldUp(crease.Position)
	default:
		log.Fatalf("'%c' is not a valid axis.", crease.Axis)
	}

	return foldedPaper
}

func (paper *Paper) foldUp(y int) *Paper {
	foldedPaper := NewPaper(y, paper.Width())

	for _, point := range paper.GetMarkedPoints() {

		if point.Y() > y {
			newY := point.Y() - (2 * common.MaxInt(0, point.Y()-y))
			foldedPoint := common.NewPoint(point.X(), newY)

			log.Tracef("Folding point (%d, %d) into (%d, %d). ", point.X(), point.Y(), foldedPoint.X(), foldedPoint.Y())

			mark := paper.GetMarkAt(point)
			if mark == '#' {
				foldedPaper.DrawMark(foldedPoint)
			}
		} else if point.Y() < y {
			log.Tracef("Point (%d, %d) doesn't need to be folded.", point.X(), point.Y())

			mark := paper.GetMarkAt(point)
			if mark == '#' {
				foldedPaper.DrawMark(point)
			}
		} else {
			log.Tracef("Point (%d, %d) is on the fold line. Skipping.", point.X(), point.Y())
		}
	}

	return foldedPaper
}

func (paper *Paper) foldLeft(x int) *Paper {
	foldedPaper := NewPaper(paper.Height(), x)

	for _, point := range paper.GetMarkedPoints() {

		if point.X() > x {
			newX := point.X() - (2 * common.MaxInt(0, point.X()-x))
			foldedPoint := common.NewPoint(newX, point.Y())

			log.Tracef("Folding point (%d, %d) into (%d, %d). ", point.X(), point.Y(), foldedPoint.X(), foldedPoint.Y())

			mark := paper.GetMarkAt(point)
			if mark == '#' {
				foldedPaper.DrawMark(foldedPoint)
			}
		} else if point.X() < x {
			log.Tracef("Point (%d, %d) doesn't need to be folded.", point.X(), point.Y())

			mark := paper.GetMarkAt(point)
			if mark == '#' {
				foldedPaper.DrawMark(point)
			}
		} else {
			log.Tracef("Point (%d, %d) is on the fold line. Skipping.", point.X(), point.Y())
		}
	}

	return foldedPaper
}

func (paper *Paper) Render() string {
	output := "Paper\n"

	for y := 0; y < paper.Height(); y++ {
		for x := 0; x < paper.Width(); x++ {
			point := common.NewPoint(x, y)
			output += string(paper.GetMarkAt(point))
		}
		output += "\n"
	}

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
