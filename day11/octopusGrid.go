package day11

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type octopusGrid struct {
	plane *common.BoundedPlane
}

func NewOctopusGrid(octopuses [][]*octopus) *octopusGrid {
	plane := common.NewBoundedPlane(len(octopuses), len(octopuses[0]))

	for row := range octopuses {
		for col := range octopuses[row] {
			plane.SetValueAt(common.New2DPoint(col, row), octopuses[row][col])
		}
	}

	return &octopusGrid{
		plane: plane,
	}
}

func (grid *octopusGrid) GetOctopusAt(point *common.Point) *octopus {
	return grid.plane.GetValueAt(point).(*octopus)
}

func (grid *octopusGrid) Step(step int) int {
	grid.charge(step)
	grid.flash(step)
	flashCount := grid.count(step)
	grid.reset()
	return flashCount
}

func (grid *octopusGrid) charge(step int) {
	log.Debug("Charging optopuses.")
	for _, point := range grid.plane.GetAllPoints() {
		octopus := grid.GetOctopusAt(point)
		octopus.Charge()
	}
}

func (grid *octopusGrid) flash(step int) {
	log.Debugf("Checking energy levels for flashes.")
	for _, point := range grid.plane.GetAllPoints() {
		octopus := grid.GetOctopusAt(point)

		if octopus.FlashIfAble(step) {
			log.Tracef("Octopus at %s has flashed. Boosting neighbors.", point)
			grid.propagateFlashFrom(point, step)
		}
	}
}

func (grid *octopusGrid) propagateFlashFrom(point *common.Point, step int) {
	adjacentPoints := grid.plane.GetMooreNeighbors(point)

	for _, adjacentPoint := range adjacentPoints {
		octopus := grid.GetOctopusAt(adjacentPoint)
		octopus.Charge()

		if octopus.FlashIfAble(step) {
			log.Tracef("Octopus at %s has flashed. Boosting neighbors.", adjacentPoint)
			grid.propagateFlashFrom(adjacentPoint, step)
		}
	}
}

func (grid *octopusGrid) count(step int) int {
	log.Debug("Counting flashes.")
	flashCount := 0

	for _, point := range grid.plane.GetAllPoints() {
		octopus := grid.GetOctopusAt(point)

		if octopus.IsFlashing(step) {
			flashCount++
		}
	}

	return flashCount
}

func (grid *octopusGrid) reset() {
	log.Debug("Reseting energy levels of optopuses that flashed.")
	for _, point := range grid.plane.GetAllPoints() {
		octopus := grid.GetOctopusAt(point)
		octopus.StopFlashing()
	}
}

func (grid *octopusGrid) Render() string {
	output := "Octopus Grid\n"

	for y := grid.plane.Span().Start().Y(); y <= grid.plane.Span().End().Y(); y++ {
		for x := grid.plane.Span().Start().X(); x <= grid.plane.Span().End().X(); x++ {
			octopus := grid.GetOctopusAt(common.New2DPoint(x, y))
			output += strconv.Itoa(octopus.energyLevel % 10)
		}
		output += "\n"
	}

	return output
}
