package day15

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	caveMap := parseCaveMap(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(caveMap),
		Part2: solvePart2(caveMap),
	}
}

func solvePart1(caveMap *CaveMap) string {
	log.Info("Solving part 1.")

	log.Info("Cave Map\n" + caveMap.String())
	optimalPath := caveMap.FindLowestRiskPath()
	log.Infof("optimalPath = %s", optimalPath)
	riskLevel := caveMap.GetRiskLevelOf(optimalPath)
	log.Infof("riskLevel = %d", riskLevel)

	log.Info("Part 1 solved.")
	return strconv.Itoa(riskLevel)
}

func solvePart2(caveMap *CaveMap) string {
	log.Info("Solving part 2.")

	caveMap = caveMap.Expand(5)
	log.Info("Cave Map\n" + caveMap.String())
	optimalPath := caveMap.FindLowestRiskPath()
	log.Infof("optimalPath = %s", optimalPath)
	riskLevel := caveMap.GetRiskLevelOf(optimalPath)
	log.Infof("riskLevel = %d", riskLevel)

	log.Info("Part 2 solved.")
	return strconv.Itoa(riskLevel)
}

func parseCaveMap(text string) *CaveMap {
	lines := common.Split(text, "\n")
	riskLevels := [][]int{}

	for row, line := range lines {
		if line == "" {
			continue
		}

		riskLevels = append(riskLevels, make([]int, len(line)))

		for col, character := range line {
			value, err := strconv.Atoi(string(character))

			if err != nil {
				log.Fatalf("'%c' is not an integer.", character)
			}

			riskLevels[row][col] = value
		}
	}

	caveMap := NewCaveMapFromValues(riskLevels)
	return caveMap
}
