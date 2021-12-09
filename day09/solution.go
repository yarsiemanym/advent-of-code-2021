package day09

import (
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	heightMap := parseHeightMap(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(heightMap),
		Part2: solvePart2(heightMap),
	}
}

func solvePart1(heightMap *common.BoundedPlane) string {
	log.Info("Solving part 1.")

	lowPoints := findLowPoints(heightMap)
	log.Debugf("%v low points found.", len(lowPoints))
	totalRiskLevel := 0

	for index, lowPoint := range lowPoints {
		log.Debugf("Assessing risk level of lowpoint %v.", index)
		log.Tracef("lowPoint = %v", *lowPoint)
		risklevel := 1 + heightMap.GetValueAt(lowPoint).(int)
		log.Tracef("risklevel = %v", risklevel)
		totalRiskLevel += risklevel
	}

	log.Tracef("totalRiskLevel = %v", totalRiskLevel)

	log.Info("Part 1 solved.")
	return strconv.Itoa(totalRiskLevel)
}

func findLowPoints(heightMap *common.BoundedPlane) []*common.Point {
	var lowPoints []*common.Point

	for _, point := range heightMap.GetAllPoints() {
		if isLowPoint(point, heightMap) {
			lowPoints = append(lowPoints, point)
		}
	}

	return lowPoints
}

func isLowPoint(point *common.Point, heightMap *common.BoundedPlane) bool {
	currentValue := heightMap.GetValueAt(point).(int)
	adjacentPoints := heightMap.GetPointsAdjacentTo(point)
	isLow := true

	for _, adjacentPoint := range adjacentPoints {
		adjacentValue := heightMap.GetValueAt(adjacentPoint).(int)
		if adjacentValue <= currentValue {
			isLow = false
			break
		}
	}

	return isLow
}

func solvePart2(heightMap *common.BoundedPlane) string {
	log.Info("Solving part 2.")

	lowPoints := findLowPoints(heightMap)
	basinSizes := make([]int, len(lowPoints))

	for index, lowPoint := range lowPoints {
		basinSizes[index] = len(exploreBasin(lowPoint, heightMap, make([]*common.Point, 0)))
	}

	sort.Ints(basinSizes)
	length := len(basinSizes)

	log.Tracef("basinSizes = %v", basinSizes)
	product := basinSizes[length-1] * basinSizes[length-2] * basinSizes[length-3]
	log.Tracef("product = %v", product)

	log.Info("Part 2 solved.")
	return strconv.Itoa(product)
}

func exploreBasin(lowPoint *common.Point, heightMap *common.BoundedPlane, alreadyExplored []*common.Point) []*common.Point {
	log.Debugf("Looking for paths uphill from point %v.", *lowPoint)
	currentValue := heightMap.GetValueAt(lowPoint).(int)
	adjacentPoints := heightMap.GetPointsAdjacentTo(lowPoint)
	pointsInBasin := []*common.Point{
		lowPoint,
	}

	for _, adjacentPoint := range adjacentPoints {
		if sliceContainsPoint(alreadyExplored, adjacentPoint) {
			log.Tracef("Already explored point %v. Skipping.", *adjacentPoint)
			continue
		}

		adjacentValue := heightMap.GetValueAt(adjacentPoint).(int)

		if adjacentValue < 9 && adjacentValue >= currentValue {
			log.Tracef("Point %v is uphill. Climbing.", *adjacentPoint)
			uphillPoints := exploreBasin(adjacentPoint, heightMap, pointsInBasin)

			for _, uphillPoint := range uphillPoints {
				if !sliceContainsPoint(pointsInBasin, uphillPoint) {
					pointsInBasin = append(pointsInBasin, uphillPoint)
				}
			}
		}
	}

	return pointsInBasin
}

func sliceContainsPoint(slice []*common.Point, element *common.Point) bool {
	for _, member := range slice {
		if *member == *element {
			return true
		}
	}

	return false
}

func parseHeightMap(text string) *common.BoundedPlane {
	lines := common.Split(text, "\n")
	var heightMap [][]int

	for row, line := range lines {
		if line == "" {
			continue
		}

		heightMap = append(heightMap, make([]int, len(line)))

		for col, char := range line {
			height, err := strconv.Atoi(string(char))

			if err != nil {
				log.Fatalf("'%c' is not an integer.", char)
			}

			heightMap[row][col] = height
		}
	}

	plane := common.NewBoundedPlane(len(heightMap), len(heightMap[0]))

	for row := range heightMap {
		for col := range heightMap[row] {
			plane.SetValueAt(common.NewPoint(col, row), heightMap[row][col])
		}
	}

	return plane
}
