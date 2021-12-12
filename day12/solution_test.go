package day12

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       12,
		InputFile: "test1.txt",
	}

	expectedPart1 := "10"
	expectedPart2 := "36"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test2(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       12,
		InputFile: "test2.txt",
	}

	expectedPart1 := "19"
	expectedPart2 := "103"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test3(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       12,
		InputFile: "test3.txt",
	}

	expectedPart1 := "226"
	expectedPart2 := "3509"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}
