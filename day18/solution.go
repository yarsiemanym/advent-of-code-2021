package day18

import (
	"math"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseSnailfishNumber)
	snailfishNumbers := make([]*SnailfishNumber, len(results))

	for index, result := range results {
		snailfishNumbers[index] = result.(*SnailfishNumber)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(snailfishNumbers),
		Part2: solvePart2(snailfishNumbers),
	}
}

func solvePart1(snailfishNumbers []*SnailfishNumber) string {
	log.Info("Solving part 1.")

	sum := snailfishNumbers[0]

	for i := 1; i < len(snailfishNumbers); i++ {
		sum = sum.Add(snailfishNumbers[i])
	}

	log.Infof("Sum is \"%s\".", sum)
	magnitude := sum.GetValue()

	log.Info("Part 1 solved.")
	return strconv.Itoa(magnitude)
}

func solvePart2(snailfishNumbers []*SnailfishNumber) string {
	log.Info("Solving part 2.")

	largestMagnitude := math.MinInt

	for i := 0; i < len(snailfishNumbers); i++ {
		for j := 0; j < len(snailfishNumbers); j++ {
			if i == j {
				continue
			}

			number1 := snailfishNumbers[i]
			number2 := snailfishNumbers[j]
			
			magnitude := number1.Add(number2).GetValue()

			if magnitude > largestMagnitude {
				largestMagnitude = magnitude
			}

			magnitude = number2.Add(number1).GetValue()

			if magnitude > largestMagnitude {
				largestMagnitude = magnitude
			}
		}
	}

	log.Info("Part 2 solved.")
	return strconv.Itoa(largestMagnitude)
}

func parseSnailfishNumber(text string) interface{} {
	if text == "" {
		return nil
	}

	return NewSnailfishNumber(text)
}
