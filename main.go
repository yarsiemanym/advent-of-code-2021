package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mitchellh/go-wordwrap"
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
	"github.com/yarsiemanym/advent-of-code-2021/day00"
)

func main() {
	common.InitLogging()
	common.InitSession()

	arg1 := os.Args[1]
	day, err := strconv.Atoi(arg1)

	if arg1 == "-h" || arg1 == "help" {
		printUsage()
		return
	}

	if err != nil {
		log.Fatalf("'%v' is not an integer", arg1)
	}

	fmt.Printf("Running day %v.\n", day)

	var input string

	switch day {
	case 0:
		log.Info("Ensuring input file exists ...")
		input = common.EnsureInputExists(day)
		log.Info("Input file exists.")
	default:
		log.Fatalf("Day %v has not been implemented.", day)
	}

	var answerPart1, answerPart2 string

	log.Info("Solving puzzle ...")

	switch day {
	case 0:
		answerPart1, answerPart2 = day00.Solve(input)
	default:
		log.Fatalf("Day %v has not been implemented.", day)
	}

	log.Info("Puzzle solved!")

	fmt.Println("Answers")
	fmt.Println("----------")
	fmt.Printf("Part 1: %v\n", answerPart1)
	fmt.Printf("Part 2: %v\n", answerPart2)
}

func printUsage() {
	limit := uint(100)
	description := "Run the solution for the puzzle from specified day of Advent of Code 2021. If a local copy of your puzzle input does not exist, it will attempt to automatically download it using your session token. When complete, the answers to parts 1 and 2 will be printed to the terminal."

	fmt.Println(wordwrap.WrapString(description, limit))
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tadvent-of-code-2021 <day>")
	fmt.Println("")
	fmt.Println("Parameters:")
	fmt.Println("")
	fmt.Println("\tday\t\t\tRun the solution for the specified day, i.e. 1-25.")
	fmt.Println("")
	fmt.Println("Environment Variables:")
	fmt.Println("")
	fmt.Println("\tAOC_SESSION_TOKEN\tSet your Advent of Code session token.")
	fmt.Println("\tAOC_LOG_LEVEL\t\tSet the log level. Defaults to 'warn' if not set.")
	fmt.Println("")
}
