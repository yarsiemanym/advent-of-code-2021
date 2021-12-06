package day06

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, ",", parseLanternfish)
	population := make([]*lanternfish, len(results))

	for i, result := range results {
		population[i] = result.(*lanternfish)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(population),
		Part2: solvePart2(population),
	}
}

func solvePart1(population []*lanternfish) string {
	log.Info("Solving part 1.")

	count := runSimulation(population, 80)

	log.Info("Part 1 solved.")
	return strconv.Itoa(count)
}

func solvePart2(population []*lanternfish) string {
	log.Info("Solving part 2.")

	// TODO: Does not finish in a reasonable amount of time.
	count := runSimulation(population, 256)

	log.Info("Part 2 solved.")
	return strconv.Itoa(count)
}

func runSimulation(population []*lanternfish, days int) int {
	finalPopulationSize := 0

	for index, fish := range population {
		log.Debugf("Counting ancestors of fish %v.", index)
		finalPopulationSize += fish.AncestorsAfter(days)
	}

	log.Tracef("finalPopulationSize = %v", finalPopulationSize)

	return finalPopulationSize
}

func parseLanternfish(text string) interface{} {
	text = strings.Trim(text, "\n")

	if text == "" {
		return nil
	}

	timer, err := strconv.Atoi(text)

	if err != nil {
		log.Fatalf("\"%v\" is not an integer.", text)
	}

	fish := &lanternfish{}
	fish.Init(timer)

	return fish
}
