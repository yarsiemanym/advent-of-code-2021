package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mitchellh/go-wordwrap"
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
	"github.com/yarsiemanym/advent-of-code-2021/day00"
	"github.com/yarsiemanym/advent-of-code-2021/day01"
	"github.com/yarsiemanym/advent-of-code-2021/day02"
)

func main() {
	common.InitLogging()
	common.InitSession()
	checkForHelpCommand()

	puzzle := setupPuzzle()
	answer := puzzle.Solve()
	fmt.Printf("Part 1 Answer: %v\n", answer.Part1)
	fmt.Printf("Part 2 Answer: %v\n", answer.Part2)
}

func checkForHelpCommand() {
	arg1 := os.Args[1]

	if arg1 == "-h" || arg1 == "--help" || arg1 == "help" {
		printUsage()
		os.Exit(0)
	}
}

func printUsage() {
	limit := uint(100)
	description1 := "Run the solution for the puzzle from specified day of Advent of Code 2021. If a local copy of your puzzle input does not exist, it will attempt to automatically download it using your session token. When complete, the answers to parts 1 and 2 will be printed to the terminal."
	description2 := "Day 0 is a special day with a mock-puzzle to exercise the application before the first real puzzle unlocks."

	fmt.Println(wordwrap.WrapString(description1, limit))
	fmt.Println("")
	fmt.Println(wordwrap.WrapString(description2, limit))
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tadvent-of-code-2021 <day>")
	fmt.Println("")
	fmt.Println("Parameters:")
	fmt.Println("")
	fmt.Println("\tday\t\t\tRun the solution for the specified day, i.e. 0-25.")
	fmt.Println("")
	fmt.Println("Environment Variables:")
	fmt.Println("")
	fmt.Println("\tAOC_SESSION_TOKEN\tSet your Advent of Code session token.")
	fmt.Println("\tAOC_LOG_LEVEL\t\tSet the log level. Defaults to \"warn\" if not set.")
	fmt.Println("")
}

func setupPuzzle() common.Puzzle {
	log.Debug("Setting up puzzle.")
	day := sanitizeDayArg(os.Args[1])

	puzzle := common.Puzzle{
		Year: 2021,
		Day:  day,
	}

	switch puzzle.Day {
	case 0:
		puzzle.SetSolution(day00.Solve)
	case 1:
		puzzle.SetSolution(day01.Solve)
	case 2:
		puzzle.SetSolution(day02.Solve)
	default:
		log.Fatalf("Day %v has no solution yet.", puzzle.Day)
	}

	return puzzle
}

func sanitizeDayArg(arg string) int {
	log.Debug("Sanitizing day argument.")
	log.Tracef("arg = \"%v\"", arg)

	day, err := strconv.Atoi(arg)

	if err != nil {
		log.Fatalf("\"%v\" is not an integer.", arg)
	}

	if day < 0 || day > 25 {
		log.Fatalf("%v is not between 0 and 25.", arg)
	}

	log.Tracef("day = %v", day)
	return day
}
