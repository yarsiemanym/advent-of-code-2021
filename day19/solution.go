package day19

import (
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n\n", parseScanner)
	scanners := make([]*Scanner, len(results))

	for index, result := range results {
		scanners[index] = result.(*Scanner)
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(scanners),
		Part2: solvePart2(scanners),
	}
}

func solvePart1(scanners []*Scanner) string {
	log.Info("Solving part 1.")

	globalMap := scanners[0]
	unknownScanners := scanners[1:]

	for len(unknownScanners) > 0 {
		unknownScanner := unknownScanners[0]
		log.Debugf("Analyzing scanner %d.", unknownScanner.Id())
		overplaps := globalMap.DetectOverlap(unknownScanner)

		if !overplaps {
			log.Debug("Scanner does not overlap known region.")
			unknownScanners = append(unknownScanners[1:], unknownScanners[0])
			continue
		}

		aligned, unknownScanner, difference := globalMap.Align(unknownScanner)

		if !aligned {
			log.Debugf("Could not align scanner %d with the scanner %d.", unknownScanner.Id(), globalMap.Id())
			unknownScanners = append(unknownScanners[1:], unknownScanners[0])
			continue
		}

		globalMap = globalMap.Merge(unknownScanner, difference)
		unknownScanners = unknownScanners[1:]
	}

	beaconsCount := len(globalMap.Beacons())

	log.Info("Part 1 solved.")
	return strconv.Itoa(beaconsCount)
}

func solvePart2(scanners []*Scanner) string {
	log.Info("Solving part 2.")

	globalMap := scanners[0]
	unknownScanners := scanners[1:]
	knownScanners := map[int]*common.Point{}

	for len(unknownScanners) > 0 {
		unknownScanner := unknownScanners[0]
		log.Debugf("Analyzing scanner %d.", unknownScanner.Id())
		overplaps := globalMap.DetectOverlap(unknownScanner)

		if !overplaps {
			log.Debug("Scanner does not overlap known region.")
			unknownScanners = append(unknownScanners[1:], unknownScanners[0])
			continue
		}

		aligned, unknownScanner, difference := globalMap.Align(unknownScanner)

		if !aligned {
			log.Debugf("Could not align scanner %d with the scanner %d.", unknownScanner.Id(), globalMap.Id())
			unknownScanners = append(unknownScanners[1:], unknownScanners[0])
			continue
		}

		globalMap = globalMap.Merge(unknownScanner, difference)
		knownScanners[unknownScanner.Id()] = difference
		unknownScanners = unknownScanners[1:]
	}

	farthestDistance := 0

	for _, position1 := range knownScanners {
		for _, position2 := range knownScanners {
			distance := position1.ManhattanDistance(position2)

			if distance > farthestDistance {
				farthestDistance = distance
			}
		}
	}

	log.Info("Part 2 solved.")
	return strconv.Itoa(farthestDistance)
}

func parseScanner(text string) interface{} {
	idPattern := regexp.MustCompile(`--- scanner (\d+) ---`)
	matches := idPattern.FindStringSubmatch(text)

	if len(matches) != 2 {
		log.Fatal("Could not extract scanner Id.")
	}

	id, err := strconv.Atoi(matches[1])

	if err != nil {
		log.Fatalf("\"%s\" is not a valid integer.", matches[1])
	}

	text = strings.Trim(text, " \n")
	lines := common.Split(text, "\n")[1:]
	pointPattern := regexp.MustCompile(`(-?\d+),(-?\d+),(-?\d+)`)
	beacons := make([]*common.Point, len(lines))

	for index, line := range lines {
		matches := pointPattern.FindStringSubmatch(line)

		if len(matches) != 4 {
			log.Fatalf("\"%s\" is not a valid point.", line)
		}

		x, err := strconv.Atoi(matches[1])

		if err != nil {
			log.Fatalf("\"%s\" is not a valid integer.", matches[1])
		}

		y, err := strconv.Atoi(matches[2])

		if err != nil {
			log.Fatalf("\"%s\" is not a valid integer.", matches[2])
		}

		z, err := strconv.Atoi(matches[3])

		if err != nil {
			log.Fatalf("\"%s\" is not a valid integer.", matches[3])
		}

		beacons[index] = common.New3DPoint(x, y, z)
	}

	return NewScanner(id, beacons)
}
