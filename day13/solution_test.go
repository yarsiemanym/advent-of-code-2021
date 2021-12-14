package day13

import (
	"strings"
	"testing"

	"github.com/yarsiemanym/advent-of-code-2021/common"
	"github.com/yarsiemanym/advent-of-code-2021/vt100"
)

var yellowBlock = vt100.Sprint(" ", vt100.YellowBackgroundAttribute)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2021,
		Day:       13,
		InputFile: "test1.txt",
	}

	expectedPart1 := "17"
	expectedPart2 := "\n" +
		"#####\n" +
		"#...#\n" +
		"#...#\n" +
		"#...#\n" +
		"#####\n" +
		".....\n" +
		"....."
	expectedPart2 = strings.Replace(expectedPart2, "#", yellowBlock, -1)
	expectedPart2 = strings.Replace(expectedPart2, ".", " ", -1)

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedPart2, answer.Part2)
	}
}
