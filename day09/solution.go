package day09

import (
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	// TODO: parse input

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(),
		Part2: solvePart2(),
	}
}

func solvePart1() string {
	log.Info("Solving part 1.")

	// TODO: implement part 1 solution

	log.Info("Part 1 solved.")
	return "Not implemented."
}

func solvePart2() string {
	log.Info("Solving part 2.")

	// TODO: implement part 2 solution

	log.Info("Part 2 solved.")
	return "Not implemented."
}
