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

func solvePart1(heightMap *heightMap) string {
	log.Info("Solving part 1.")

	lowPoints := heightMap.FindLowPoints()
	log.Debugf("%v low points found.", len(lowPoints))
	log.Info(heightMap.RenderLowPoints(lowPoints))
	totalRiskLevel := 0

	for index, lowPoint := range lowPoints {
		log.Debugf("Assessing risk level of low point %v.", index)
		log.Tracef("lowPoint = %v", *lowPoint)
		risklevel := 1 + heightMap.GetHeightAt(lowPoint)
		log.Tracef("risklevel = %v", risklevel)
		totalRiskLevel += risklevel
	}

	log.Tracef("totalRiskLevel = %v", totalRiskLevel)

	log.Info("Part 1 solved.")
	return strconv.Itoa(totalRiskLevel)
}

func solvePart2(heightMap *heightMap) string {
	log.Info("Solving part 2.")

	lowPoints := heightMap.FindLowPoints()
	basins := make([]*basin, len(lowPoints))
	basinSizes := make([]int, len(lowPoints))

	for index, lowPoint := range lowPoints {
		basins[index] = heightMap.ExploreBasin(lowPoint, NewBasin())
		basinSizes[index] = basins[index].Size()
	}

	log.Debugf("%v basins found.", len(basins))
	log.Info(heightMap.RenderBasins(basins))

	sort.Ints(basinSizes)
	length := len(basinSizes)

	log.Tracef("basinSizes = %v", basinSizes)
	productOfLargestThreeBasinSizes := basinSizes[length-1] * basinSizes[length-2] * basinSizes[length-3]
	log.Tracef("productOfLargestThreeBasinSizes = %v", productOfLargestThreeBasinSizes)

	log.Info("Part 2 solved.")
	return strconv.Itoa(productOfLargestThreeBasinSizes)
}

func parseHeightMap(text string) *heightMap {
	lines := common.Split(text, "\n")
	var heights [][]int

	for row, line := range lines {
		if line == "" {
			continue
		}

		heights = append(heights, make([]int, len(line)))

		for col, char := range line {
			height, err := strconv.Atoi(string(char))

			if err != nil {
				log.Fatalf("'%c' is not an integer.", char)
			}

			heights[row][col] = height
		}
	}

	return NewHeightMap(heights)
}
