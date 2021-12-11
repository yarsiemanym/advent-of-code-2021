package day09

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type heightMap struct {
	plane *common.BoundedPlane
}

func NewHeightMap(heights [][]int) *heightMap {
	plane := common.NewBoundedPlane(len(heights), len(heights[0]))

	for row := range heights {
		for col := range heights[row] {
			plane.SetValueAt(common.NewPoint(col, row), heights[row][col])
		}
	}

	return &heightMap{
		plane: plane,
	}
}

func (heightMap *heightMap) FindLowPoints() []*common.Point {
	log.Debug("Seaching for low points.")
	var lowPoints []*common.Point

	for _, point := range heightMap.plane.GetAllPoints() {
		if heightMap.IsLowPoint(point) {
			lowPoints = append(lowPoints, point)
		}
	}

	log.Debugf("%v low points found.", len(lowPoints))

	return lowPoints
}

func (heightMap *heightMap) IsLowPoint(point *common.Point) bool {
	log.Tracef("Determining if point %v is a low point.", *point)
	currentValue := heightMap.GetHeightAt(point)
	log.Tracef("height = %v", currentValue)
	adjacentPoints := heightMap.GetPointsAdjacentTo(point)
	isLow := true

	for _, adjacentPoint := range adjacentPoints {
		log.Tracef("Inspecting adjacent point %v.", *adjacentPoint)
		adjacentValue := heightMap.GetHeightAt(adjacentPoint)
		log.Tracef("height = %v", adjacentValue)
		if adjacentValue <= currentValue {
			log.Tracef("Adjacent point %v is lower.", *adjacentPoint)
			isLow = false
			break
		} else {
			log.Tracef("Adjacent point %v is not lower.", *adjacentPoint)
		}
	}

	if isLow {
		log.Tracef("Point %v is a low point.", *point)
	} else {
		log.Tracef("Point %v is not a low point.", *point)
	}

	return isLow
}

func (heightMap *heightMap) GetHeightAt(point *common.Point) int {
	value := heightMap.plane.GetValueAt(point)

	if value == nil {
		value = 0
	}

	return value.(int)
}

func (heightMap *heightMap) GetPointsAdjacentTo(point *common.Point) []*common.Point {
	return heightMap.plane.GetPointsAdjacentTo(point)
}

func (heightMap *heightMap) ExploreBasin(lowPoint *common.Point, exploredBasin *basin) *basin {
	log.Tracef("Looking for paths up hill from point %v.", *lowPoint)
	currentValue := heightMap.GetHeightAt(lowPoint)
	adjacentPoints := heightMap.GetPointsAdjacentTo(lowPoint)
	basin := NewBasin()
	basin.Add(lowPoint)

	for _, adjacentPoint := range adjacentPoints {
		if exploredBasin.Contains(adjacentPoint) {
			log.Tracef("Already explored point %v. Skipping.", *adjacentPoint)
			continue
		}

		adjacentValue := heightMap.GetHeightAt(adjacentPoint)

		if adjacentValue < 9 && adjacentValue >= currentValue {
			log.Tracef("Point %v is up hill. Climbing.", *adjacentPoint)
			newlyExploredBasin := heightMap.ExploreBasin(adjacentPoint, basin)

			basin.Add(newlyExploredBasin.points...)
		}
	}

	return basin
}

func (heightMap *heightMap) RenderLowPoints(points []*common.Point) string {
	output := "Low Points\n"

	for y := heightMap.plane.Span().Start().Y(); y <= heightMap.plane.Span().End().Y(); y++ {
		for x := heightMap.plane.Span().Start().X(); x <= heightMap.plane.Span().End().X(); x++ {
			isLowPoint := false

			for _, point := range points {
				if point.X() == x && point.Y() == y {
					isLowPoint = true
					break
				}
			}

			if isLowPoint {
				output += strconv.Itoa(heightMap.GetHeightAt(common.NewPoint(x, y)))
			} else {
				output += "."
			}
		}

		output += "\n"
	}

	return output
}

func (heightMap *heightMap) RenderBasins(basins []*basin) string {
	output := "Basins\n"

	for y := heightMap.plane.Span().Start().Y(); y <= heightMap.plane.Span().End().Y(); y++ {
		for x := heightMap.plane.Span().Start().X(); x <= heightMap.plane.Span().End().X(); x++ {
			point := common.NewPoint(x, y)
			pointIsInABasin := false

			for _, basin := range basins {
				if basin.Contains(point) {
					pointIsInABasin = true
					break
				}
			}

			if pointIsInABasin {
				output += strconv.Itoa(heightMap.GetHeightAt(point))
			} else {
				output += "."
			}
		}

		output += "\n"
	}

	return output
}
