package day11

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	grid1 := parseOctopusGrid(text)
	grid2 := parseOctopusGrid(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(grid1),
		Part2: solvePart2(grid2),
	}
}

func solvePart1(grid *octopusGrid) string {
	log.Info("Solving part 1.")

	totalFlashes := 0

	for i := 1; i <= 100; i++ {
		log.Debugf("Simulating step %d.", i)
		stepFlashes := grid.Step(i)
		log.Debugf("%d flashes during this step.", stepFlashes)
		totalFlashes += stepFlashes
		log.Debugf("%d total flashes.", totalFlashes)
		log.Trace(grid.Render())
	}

	log.Info("Part 1 solved.")
	return strconv.Itoa(totalFlashes)
}

func solvePart2(grid *octopusGrid) string {
	log.Info("Solving part 2.")

	allFlashingStep := 0

	for i := 1; allFlashingStep == 0; i++ {
		log.Debugf("Simulating step %v.", i)
		stepFlashes := grid.Step(i)
		log.Debugf("%d flashes during this step.", stepFlashes)
		if stepFlashes == 100 {
			log.Debug("All octopuses are flashing!")
			allFlashingStep = i
		} else {
			log.Debug("Not all octopuses are flashing.")
		}

		log.Trace(grid.Render())
	}

	log.Info("Part 2 solved.")
	return strconv.Itoa(allFlashingStep)
}

func parseOctopusGrid(text string) *octopusGrid {
	octopuses := parseOctopuses(text)
	return NewOctopusGrid(octopuses)
}

func parseOctopuses(text string) [][]*octopus {
	lines := common.Split(text, "\n")
	octopuses := [][]*octopus{}

	for y, line := range lines {
		if line == "" {
			continue
		}

		octopuses = append(octopuses, make([]*octopus, len(line)))

		for x, character := range line {
			energyLevel, err := strconv.Atoi(string(character))

			if err != nil {
				log.Fatalf("'%c' is not an integer.", character)
			}

			octopuses[y][x] = NewOctopus(energyLevel)
		}
	}

	return octopuses
}
