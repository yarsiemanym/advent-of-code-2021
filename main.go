package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2021/common"
	"github.com/yarsiemanym/advent-of-code-2021/day00"
)

func main() {
	common.InitLogging()
	common.InitSession()

	day, err := strconv.Atoi(os.Args[1])
	common.Check(err)

	input := ensureInput(day)

	var answerPart1, answerPart2 string

	switch day {
	case 0:
		answerPart1, answerPart2 = day00.Solve(input)
	default:
		panic("Day %v is not implemented.")
	}

	fmt.Printf("Part 1: %v\n", answerPart1)
	fmt.Printf("Part 2: %v\n", answerPart2)
}
