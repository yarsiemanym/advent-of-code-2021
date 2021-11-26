package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mitchellh/go-wordwrap"
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
	"github.com/yarsiemanym/advent-of-code-2021/day00"
)

func main() {
	common.InitLogging()
	common.InitSession()

	arg1 := os.Args[1]

	if arg1 == "-h" || arg1 == "help" {
		printUsage()
		return
	}

	day := SanitizeDayArg(arg1)

	fmt.Printf("Running day %v solution.\n", day)

	var input string

	if IsPuzzleUnlocked(day) {
		log.Info("Ensuring input file exists ...")
		input = common.EnsureInputExists(2021, day)
		log.Info("Input file exists.")
	} else {
		log.Fatalf("Day %v has not been unlocked.", day)
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

	fmt.Println("")
	fmt.Println("Answers")
	fmt.Println("----------")
	fmt.Printf("Part 1: %v\n", answerPart1)
	fmt.Printf("Part 2: %v\n", answerPart2)
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

func SanitizeDayArg(arg string) int {
	log.Debug("Sanitizing day argument.")
	log.Tracef("arg = \"%v\"", arg)

	day, err := strconv.Atoi(arg)

	if err != nil {
		log.Fatalf("'%v' is not an integer.", arg)
	}

	if day < 0 || day > 25 {
		log.Fatalf("%v is not between 0 and 25.", arg)
	}

	return day
}

func IsPuzzleUnlocked(day int) bool {
	log.Debugf("Checking if day %v has been unlocked.", day)

	est, err := time.LoadLocation(("EST"))
	common.Check(err)

	var puzzleUnlockAt time.Time

	if day != 0 {
		puzzleUnlockAt = time.Date(2021, 11, 30, 0, 0, 0, 0, est).Add(time.Hour * 24 * time.Duration(day))
	}

	log.Tracef("puzzleUnlockAt = \"%v\"", puzzleUnlockAt)

	isUnlocked := puzzleUnlockAt.Before(time.Now())

	log.Tracef("isUnlocked = %v", isUnlocked)

	return isUnlocked
}
