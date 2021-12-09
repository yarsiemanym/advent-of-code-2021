package day09

import (
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
	var lowPoints []*common.Point

	for _, point := range heightMap.plane.GetAllPoints() {
		if heightMap.IsLowPoint(point) {
			lowPoints = append(lowPoints, point)
		}
	}

	return lowPoints
}

func (heightMap *heightMap) IsLowPoint(point *common.Point) bool {
	currentValue := heightMap.GetHeightAt(point)
	adjacentPoints := heightMap.GetPointsAdjacentTo(point)
	isLow := true

	for _, adjacentPoint := range adjacentPoints {
		adjacentValue := heightMap.GetHeightAt(adjacentPoint)
		if adjacentValue <= currentValue {
			isLow = false
			break
		}
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

func (heightMap *heightMap) ExploreBasin(lowPoint *common.Point, alreadyExplored []*common.Point) []*common.Point {
	log.Debugf("Looking for paths uphill from point %v.", *lowPoint)
	currentValue := heightMap.GetHeightAt(lowPoint)
	adjacentPoints := heightMap.GetPointsAdjacentTo(lowPoint)
	pointsInBasin := []*common.Point{
		lowPoint,
	}

	for _, adjacentPoint := range adjacentPoints {
		if sliceContainsPoint(alreadyExplored, adjacentPoint) {
			log.Tracef("Already explored point %v. Skipping.", *adjacentPoint)
			continue
		}

		adjacentValue := heightMap.GetHeightAt(adjacentPoint)

		if adjacentValue < 9 && adjacentValue >= currentValue {
			log.Tracef("Point %v is uphill. Climbing.", *adjacentPoint)
			uphillPoints := heightMap.ExploreBasin(adjacentPoint, pointsInBasin)

			for _, uphillPoint := range uphillPoints {
				if !sliceContainsPoint(pointsInBasin, uphillPoint) {
					pointsInBasin = append(pointsInBasin, uphillPoint)
				}
			}
		}
	}

	return pointsInBasin
}
