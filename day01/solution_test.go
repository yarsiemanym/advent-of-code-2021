package day01

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := common.Puzzle{
		Year:      2021,
		Day:       0,
		InputFile: "test1.txt",
	}

	part1Expected := "7"
	part2Expected := "5"

	answer := Solve(puzzle)

	if answer.Part1 != part1Expected {
		t.Errorf("Part 1: Expected \"%v\" but got \"%v\".", part1Expected, answer.Part1)
	}

	if answer.Part2 != part2Expected {
		t.Errorf("Part 2: Expected \"%v\" but got \"%v\".", part2Expected, answer.Part2)
	}
}
