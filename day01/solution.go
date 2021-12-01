package day01

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parseDepth)
	measurements := make([]int, len(results))

	for index, result := range results {
		measurements[index] = cast.ToInt(result)
	}

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(measurements),
		Part2: solvePart2(measurements),
	}

	return answer
}

func solvePart1(measurements []int) string {
	log.Debug("Solving part 1.")
	log.Tracef("measurements = %v", measurements)
	previousMeasurement := 0
	increases := 0

	for index, measurement := range measurements {
		log.Debugf("Checking measurement %v.", index)
		log.Tracef("previousMeasurement = %v", previousMeasurement)
		log.Tracef("increases = %v", increases)
		log.Tracef("measurement = %v", measurement)

		if previousMeasurement != 0 && measurement > previousMeasurement {
			log.Debug("Decrease detected.")
			increases++
		}

		previousMeasurement = measurement
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(increases)
}

func solvePart2(measurements []int) string {
	log.Debug("Solving part 2.")
	log.Tracef("measurements = %v", measurements)
	previousMeasurement := 0
	increases := 0

	for i := 0; i < len(measurements)-2; i++ {
		log.Debugf("Checking measurement %v.", i)
		log.Tracef("previousMeasurement = %v", previousMeasurement)
		log.Tracef("decreases = %v", increases)

		measurement := measurements[i] + measurements[i+1] + measurements[i+2]
		log.Tracef("measurement = %v", measurement)

		if previousMeasurement != 0 && measurement > previousMeasurement {
			log.Debug("Decrease detected.")
			increases++
		}

		previousMeasurement = measurement
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(increases)
}

func parseDepth(text string) interface{} {
	if text == "" {
		return nil
	}

	result, err := strconv.Atoi(text)
	common.Check(err)

	return result
}
