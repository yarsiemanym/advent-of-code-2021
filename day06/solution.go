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
	for day := 1; day <= days; day++ {
		log.Debugf("Simulating day %v.", day)
		var babies []*lanternfish

		for index, fish := range population {
			log.Tracef("Aging lanternfish %v.", index)
			baby := fish.Tick()

			if baby != nil {
				log.Debugf("Lanternfish %v had a baby!", index)
				babies = append(babies, baby)
			}
		}

		log.Tracef("Adding %v babies to the lanternfish population.", len(babies))
		population = append(population, babies...)
		log.Debugf("At the end of day %v, the lanternfish population is %v.", day, len(population))
	}

	return len(population)
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
