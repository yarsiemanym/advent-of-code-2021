package day03

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseDiagnotic)
	diagnotics := make([]diagnostic, len(results))

	for index, result := range results {
		diagnotics[index] = result.(diagnostic)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(diagnotics),
		Part2: solvePart2(diagnotics),
	}
}

func solvePart1(diagnotics []diagnostic) string {
	log.Debug("Solving part 1.")
	log.Tracef("diagnostics = %v", diagnotics)

	gamma := 0
	epsilon := 0

	log.Debug("Inspecting bits.")

	size := diagnotics[0].Size // Assume they are all the same size.
	log.Tracef("size = %v", size)

	for position := 0; position < size; position++ {
		log.Debugf("Inspecting position %v.", position)

		mask := 1 << position
		log.Tracef("mask = %b", mask)

		if isOneMostCommon(diagnotics, mask) {
			log.Debug("0 is most common.")
			gamma &^= mask
			epsilon |= mask
		} else {
			log.Debug("1 is most common.")
			gamma |= mask
			epsilon &^= mask
		}
	}

	log.Debug("Inspection complete.")
	log.Tracef("gamma = %v", gamma)
	log.Tracef("epsilon = %v", epsilon)

	log.Debug("Part 1 solved.")
	return strconv.Itoa(gamma * epsilon)
}

func isOneMostCommon(diagnotics []diagnostic, mask int) bool {
	zeroCount := 0
	oneCount := 0

	for _, diagnostic := range diagnotics {
		maskedValue := diagnostic.Value & mask

		if maskedValue == 0 {
			zeroCount++
		} else {
			oneCount++
		}
	}

	return oneCount >= zeroCount
}

func solvePart2(diagnostics []diagnostic) string {
	log.Debug("Solving part 2.")

	oxygenGeneratorRating := determinOxygenGeneratorRating(diagnostics)
	co2ScrubberRating := determinCo2ScrubberRating(diagnostics)

	log.Debug("Part 2 solved.")
	return strconv.Itoa(oxygenGeneratorRating * co2ScrubberRating)
}

func determinOxygenGeneratorRating(diagnostics []diagnostic) int {
	log.Debugf("Determining oxygen generator rating.")
	rating := determinLifeSupportRating(diagnostics, true)
	log.Tracef("rating = %v", rating)
	return rating
}

func determinCo2ScrubberRating(diagnostics []diagnostic) int {
	log.Debugf("Determining CO2 scrubber rating.")
	rating := determinLifeSupportRating(diagnostics, false)
	log.Tracef("rating = %v", rating)
	return rating
}

func determinLifeSupportRating(diagnostics []diagnostic, keepMostCommon bool) int {
	log.Debugf("Determining life support rating.")
	log.Tracef("keepMostCommon = %v", keepMostCommon)
	size := diagnostics[0].Size // Assume they are all the same size.

	remainingDiagnostics := diagnostics
	for position := size - 1; position >= 0 && len(remainingDiagnostics) > 1; position-- {
		log.Debugf("Inspecting position %v.", position)
		remainingDiagnostics = filterDiagnostics(remainingDiagnostics, position, keepMostCommon)
	}

	return remainingDiagnostics[0].Value
}

func filterDiagnostics(diagnostics []diagnostic, position int, keepMostCommon bool) []diagnostic {
	log.Debugf("Detecting most common bit at position %v.", position)

	mask := 1 << position
	log.Tracef("inspectionMask = %b", mask)

	oneIsMostCommon := isOneMostCommon(diagnostics, mask)

	if oneIsMostCommon {
		log.Trace("1 is most common.")
	} else {
		log.Trace("0 is most common.")
	}

	log.Debugf("Filtering diagnostics at position %v.", position)
	log.Tracef("keepMostCommon = %v", keepMostCommon)

	var remainingDiagnostics []diagnostic
	for _, diagnostic := range diagnostics {
		if keepDiagnostic(diagnostic, mask, keepMostCommon, oneIsMostCommon) {
			remainingDiagnostics = append(remainingDiagnostics, diagnostic)
		}
	}

	log.Tracef("remainingDiagnostics = %v", remainingDiagnostics)
	return remainingDiagnostics
}

func keepDiagnostic(diagnostic diagnostic, mask int, keepMostCommon bool, oneIsMostCommon bool) bool {
	return ((diagnostic.Value&mask != 0) == oneIsMostCommon) == keepMostCommon
}

func parseDiagnotic(text string) interface{} {
	if text == "" {
		return nil
	}

	value, err := strconv.ParseInt(text, 2, len(text)+1)
	common.Check(err)

	diagnostic := diagnostic{
		Value: int(value),
		Size:  len(text),
	}
	return diagnostic
}
