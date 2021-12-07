package day07

import (
	"math"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, ",", parseCrabSubmarinePositions)
	positions := make([]int, len(results))

	for index, result := range results {
		positions[index] = result.(int)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(positions),
		Part2: solvePart2(positions),
	}
}

func solvePart1(positions []int) string {
	log.Info("Solving part 1.")
	minFuel := determineOptimalAlignmentPosition(positions, calulateFuelCostPart1)
	log.Info("Part 1 solved.")
	return strconv.Itoa(minFuel)
}

func solvePart2(positions []int) string {
	log.Info("Solving part 2.")
	minFuel := determineOptimalAlignmentPosition(positions, calulateFuelCostPart2)
	log.Info("Part 2 solved.")
	return strconv.Itoa(minFuel)
}

type fuelCalculator func(int) int // distance -> fuel

func determineOptimalAlignmentPosition(positions []int, fuelCalculator fuelCalculator) int {
	log.Debug("Determining optimal alignment position.")
	max := common.MaxInt(positions...)
	currentOptimumTarget := 0
	currentMinFuel := math.MaxInt

	for target := 0; target < max; target++ {
		log.Debugf("Evaluating position %v.", target)
		fuel := 0

		for _, position := range positions {
			log.Tracef("Evaluating travel from position %v to position %v.", position, target)
			distance := common.AbsInt(position - target)
			log.Tracef("distance = %v", distance)
			fuel += fuelCalculator(distance)
		}

		log.Debugf("Aligning on position %v costs %v fuel.", target, fuel)
		log.Debug("Checking for new optimum target.")
		log.Tracef("currentOptimumTarget = %v", currentOptimumTarget)
		log.Tracef("currentMinFuel = %v", currentMinFuel)

		if fuel < currentMinFuel {
			log.Debugf("Position %v is the new optimum target. Total fuel consumption is %v.", target, fuel)
			currentOptimumTarget = target
			currentMinFuel = fuel
		} else {
			log.Debugf("Position %v is still the optimum target. Total fuel consumption is %v.", currentOptimumTarget, currentMinFuel)
		}
	}

	log.Debug("Evaluation complete.")
	log.Tracef("The optimum alignment position is %v which costs %v fuel.", currentOptimumTarget, currentMinFuel)
	return currentMinFuel
}

func calulateFuelCostPart1(distance int) int {
	log.Tracef("Calculating fuel cost to travel distance %v.", distance)
	fuel := distance
	log.Tracef("Travelling distance %v costs %v fuel.", distance, fuel)
	return fuel
}

func calulateFuelCostPart2(distance int) int {
	log.Tracef("Calculating fuel cost to travel distance %v.", distance)
	fuel := 0
	for step := 1; step <= distance; step++ {
		fuel += step
	}
	log.Tracef("Travelling distance %v costs %v fuel.", distance, fuel)
	return fuel
}

func parseCrabSubmarinePositions(text string) interface{} {
	if text == "" {
		return nil
	}

	text = strings.Trim(text, "\n")
	position, err := strconv.Atoi(text)

	if err != nil {
		log.Fatalf("\"%v\" is not an integer.", text)
	}

	return position
}
