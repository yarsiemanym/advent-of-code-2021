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
		Part2: solvePart2(),
	}
}

func solvePart1(caveMap *CaveMap) string {
	log.Info("Solving part 1.")

	start := common.NewPoint(0, 0)
	end := common.NewPoint(caveMap.Width()-1, caveMap.Height()-1)

	log.Infof("start = %s", start)
	log.Infof("end = %s", end)

	initialPath := NewPath()
	initialPath.Append(start)

	optimalPath := caveMap.FindLowestRiskLevel(initialPath, end)
	log.Infof("optimalPath = %s", optimalPath)

	riskValue := caveMap.RiskLevelOf(optimalPath)
	log.Infof("riskValue = %d", riskValue)

	log.Info("Part 1 solved.")
	return strconv.Itoa(riskValue)
}

func solvePart2() string {
	log.Info("Solving part 2.")

	// TODO: implement part 2 solution

	log.Info("Part 2 solved.")
	return "Not implemented."
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
	log.Info("Cave Map\n" + caveMap.String())

	return caveMap
}
