package day15

import (
	"math"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type CaveMap struct {
	plane *common.BoundedPlane
}

func NewCaveMap(height int, width int) *CaveMap {
	return &CaveMap{
		plane: common.NewBoundedPlane(height, width),
	}
}

func NewCaveMapFromValues(riskLevels [][]int) *CaveMap {
	values := make([][]interface{}, len(riskLevels))

	for row := range riskLevels {
		values[row] = make([]interface{}, len(riskLevels[row]))

		for col := range riskLevels[row] {
			values[row][col] = riskLevels[row][col]
		}
	}

	return &CaveMap{
		plane: common.NewBoundedPlaneFromValues(values),
	}
}

func (caveMap *CaveMap) GetRiskLevelAt(point *common.Point) int {
	riskLevel := math.MaxInt
	value := caveMap.plane.GetValueAt(point)

	if value != nil {
		riskLevel = value.(int)
	}

	return riskLevel
}

func (caveMap *CaveMap) Height() int {
	return caveMap.plane.Span().End().Y() + 1
}

func (caveMap *CaveMap) Width() int {
	return caveMap.plane.Span().End().X() + 1
}

func (caveMap *CaveMap) GetPointsAdjacentTo(point *common.Point, excludePath *Path) []*common.Point {
	adjacentPoints := []*common.Point{}

	for _, adjacentPoint := range caveMap.plane.GetVonNeumannNeighbors(point) {
		exclude := false

		for _, excludedPoint := range excludePath.Points() {
			if *adjacentPoint == *excludedPoint {
				exclude = true
				break
			}
		}

		if !exclude {
			adjacentPoints = append(adjacentPoints, adjacentPoint)
		}
	}

	return adjacentPoints
}

func (caveMap *CaveMap) FindLowestRiskLevel(trunk *Path, end *common.Point) *Path {
	tip := trunk.End()
	log.Debugf("Finding lowest risk path from %s to %s.", tip, end)

	var optimalPath *Path
	for _, adjacentPoint := range caveMap.GetPointsAdjacentTo(tip, trunk) {
		log.Debugf("Inspecting adjacent point %s.", adjacentPoint)
		var branchOptimalPath *Path
		if *adjacentPoint == *end {
			log.Debug("Destination found!")
			branchOptimalPath = NewPath()
			branchOptimalPath.Append(adjacentPoint)
		} else {
			log.Debug("Destination not found.")
			newTrunk := trunk.Clone()
			newTrunk.Append(adjacentPoint)
			branchOptimalPath = caveMap.FindLowestRiskLevel(newTrunk, end)
			if branchOptimalPath != nil {
				branchOptimalPath.Prepend(adjacentPoint)
			}
		}

		if caveMap.RiskLevelOf(branchOptimalPath) < caveMap.RiskLevelOf(optimalPath) {
			optimalPath = branchOptimalPath
		}
	}

	log.Debugf("Lowest risk path from %s to %s is %s", tip, end, optimalPath)
	return optimalPath
}

func (caveMap *CaveMap) RiskLevelOf(path *Path) int {
	if path == nil {
		return math.MaxInt
	}

	riskLevel := 0

	for _, point := range path.Points() {
		riskLevel += caveMap.GetRiskLevelAt(point)
	}

	return riskLevel
}

func (caveMap *CaveMap) String() string {
	output := ""

	for y := 0; y < caveMap.Height(); y++ {
		for x := 0; x < caveMap.Width(); x++ {
			point := common.NewPoint(x, y)
			riskLevel := caveMap.GetRiskLevelAt(point)
			output += strconv.Itoa(riskLevel)
		}

		output += "\n"
	}

	return output
}
