package day15

import (
	"math"
	"strconv"

	"github.com/ahrtr/gocontainer/queue/priorityqueue"
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

func (caveMap *CaveMap) StraightPath() *Path {
	path := NewPath()

	for x, y := 0, 0; x < caveMap.Width()-1 && y < caveMap.Height()-1; {

		path.Append(common.NewPoint(x, y))
		x++
		path.Append(common.NewPoint(x, y))
		y++
	}

	return path
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

func (caveMap *CaveMap) FindLowestRiskPath() *Path {
	start := common.NewPoint(0, 0)
	end := common.NewPoint(caveMap.Width()-1, caveMap.Height()-1)
	_, prev := caveMap.dijkstra(start, end)

	path := NewPath()
	here := end
	for here != nil {
		path.Prepend(here)
		here = prev[*here]
	}

	if path.Start() != start && path.End() != end {
		log.Errorf("Path does not reach from %s to %s.", start, end)
		log.Tracef("path = %s", path)
		return nil
	}

	return path
}

func (caveMap *CaveMap) Compare(v1 interface{}, v2 interface{}) (int, error) {
	point1 := v1.(common.Point)
	point2 := v2.(common.Point)
	riskLevel1 := caveMap.GetRiskLevelAt(&point1)
	riskLevel2 := caveMap.GetRiskLevelAt(&point2)
	if riskLevel1 < riskLevel2 {
		return -1, nil
	} else if riskLevel1 > riskLevel2 {
		return 1, nil
	} else {
		return 0, nil
	}
}

var distances map[common.Point]int
var previous map[common.Point]*common.Point

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func (caveMap *CaveMap) dijkstra(start *common.Point, end *common.Point) (map[common.Point]int, map[common.Point]*common.Point) {
	log.Debug("Begin Dijkstra.")

	ceiling := caveMap.GetRiskLevelOf(caveMap.StraightPath())
	unvisited := priorityqueue.New().WithComparator(caveMap)
	distances = map[common.Point]int{}
	previous = map[common.Point]*common.Point{}

	unvisited.Add(*start)
	distances[*start] = 0

	for !unvisited.IsEmpty() {
		log.Debugf("%d points left to visit.", unvisited.Size())
		here := unvisited.Poll().(common.Point)

		for _, neighbor := range caveMap.plane.GetVonNeumannNeighbors(&here) {
			alternateDistance := distances[here] + caveMap.GetRiskLevelAt(neighbor)

			neightborDistance, exists := distances[*neighbor]

			if !exists {
				neightborDistance = math.MaxInt
			}

			if alternateDistance > ceiling {
				unvisited.Remove(*neighbor)
			} else if alternateDistance < neightborDistance {
				distances[*neighbor] = alternateDistance
				previous[*neighbor] = &here
				unvisited.Add(*neighbor)
			}
		}
	}

	log.Debug("End Dijkstra.")

	return distances, previous
}

func (caveMap *CaveMap) GetRiskLevelOf(path *Path) int {
	if path == nil {
		return math.MaxInt
	}

	riskLevel := 0

	for _, point := range path.Points()[1:] {
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

func (caveMap *CaveMap) Expand(coefficient int) *CaveMap {
	originalHeight := caveMap.Height()
	originalWidth := caveMap.Width()
	expandedHeight := originalHeight * coefficient
	expandedWidth := originalWidth * coefficient

	expandedRiskLevels := make([][]int, expandedHeight)
	for row := range expandedRiskLevels {
		expandedRiskLevels[row] = make([]int, expandedWidth)

		for col := range expandedRiskLevels[row] {
			riskLevel := caveMap.GetRiskLevelAt(common.NewPoint(col%originalWidth, row%originalHeight))
			expandedRiskLevel := (riskLevel + (row / originalHeight) + (col / originalWidth))
			if expandedRiskLevel > 9 {
				expandedRiskLevel -= 9
			}
			expandedRiskLevels[row][col] = expandedRiskLevel
		}
	}

	return NewCaveMapFromValues(expandedRiskLevels)
}
