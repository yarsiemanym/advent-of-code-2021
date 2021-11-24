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

	day, err := strconv.Atoi(os.Args[1])
	common.Check(err)

	input := os.Args[2]

	var answer string

	switch day {
	case 0:
		answer = day00.Solve(input)
	default:
		panic("Day %v is not implemented.")
	}

	fmt.Print(answer)
}
