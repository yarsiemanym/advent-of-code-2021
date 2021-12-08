package day08

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parseEntry)
	entries := make([]*entry, len(results))

	for index, result := range results {
		entries[index] = result.(*entry)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(entries),
		Part2: solvePart2(entries),
	}
}

func solvePart1(entries []*entry) string {
	log.Info("Solving part 1.")

	count := 0

	for _, entry := range entries {
		for _, outputValue := range entry.ouputValues {
			length := len(outputValue)
			switch length {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	log.Tracef("count = %v", count)

	log.Info("Part 1 solved.")
	return strconv.Itoa(count)
}

func solvePart2(entries []*entry) string {
	log.Info("Solving part 2.")

	sum := 0

	for index, entry := range entries {
		log.Debugf("Processing entry %v.", index)
		log.Tracef("entry = %v", entry)

		signalMapper := NewSignalMapper(entry.uniqueSignalPatterns)
		display := NewValueDisplay()
		log.Debugf("Setting signals for entry %v.", index)
		display.SetSignals(entry.ouputValues[0], entry.ouputValues[1], entry.ouputValues[2], entry.ouputValues[3], signalMapper)
		value := display.NumericValue()

		if value == nil {
			log.Fatalf("Dispay does not show a valid number for signals %v.", entry.ouputValues)
		} else {
			log.Debugf("Value displayed for entry %v is %v.", index, *value)
			sum += *value
		}
	}

	log.Info("Part 2 solved.")
	return strconv.Itoa(sum)
}

func parseEntry(text string) interface{} {
	if text == "" {
		return nil
	}

	parts := common.Split(text, "|")

	if len(parts) != 2 {
		log.Fatalf("Splitting line on '|' yielded %v parts.", len(parts))
	}

	uniqueSignalPatterns := common.Split(parts[0], " ")
	uniqueSignalPatterns = uniqueSignalPatterns[0 : len(uniqueSignalPatterns)-1]

	outputValues := common.Split(parts[1], " ")
	outputValues = outputValues[1:]

	entry := &entry{
		uniqueSignalPatterns: uniqueSignalPatterns,
		ouputValues:          outputValues,
	}

	return entry
}
