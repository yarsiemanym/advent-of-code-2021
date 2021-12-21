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

	knownRegion := scanners[0]
	unknownRegions := scanners[1:]

	for len(unknownRegions) > 0 {
		unknownRegion := unknownRegions[0]
		log.Debugf("Analyzing scanner %d.", unknownRegion.Id())
		overplaps := false

		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				for z := 0; z < 4; z++ {
					var difference *common.Point
					overplaps, difference = knownRegion.DetectBeaconOverlap(unknownRegion)
					if overplaps {
						log.Debug("Overlap detected! Merging beacons.")
						knownRegion = knownRegion.Merge(unknownRegion, difference)
						break
					}
					unknownRegion = unknownRegion.RotateZClockwise()

				}

				if overplaps {
					break
				}
				unknownRegion = unknownRegion.RotateYClockwise()
			}

			if overplaps {
				break
			}
			unknownRegion = unknownRegion.RotateXClockwise()
		}

		if !overplaps {
			log.Debug("Scanner does not overlap known region.")
			unknownRegions = append(unknownRegions[1:], unknownRegions[0])
		} else {
			unknownRegions = unknownRegions[1:]
		}
	}

	beaconsCount := len(knownRegion.Beacons())

	log.Info("Part 1 solved.")
	return strconv.Itoa(beaconsCount)
}

func solvePart2(scanners []*Scanner) string {
	log.Info("Solving part 2.")

	gobalMap := scanners[0]
	unknownScanners := scanners[1:]
	knownScanners := map[int]*common.Point{}

	for len(unknownScanners) > 0 {
		unknownScanner := unknownScanners[0]
		log.Debugf("Analyzing scanner %d.", unknownScanner.Id())
		overplaps := false

		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				for z := 0; z < 4; z++ {
					var difference *common.Point
					overplaps, difference = gobalMap.DetectBeaconOverlap(unknownScanner)
					if overplaps {
						log.Debug("Overlap detected! Merging beacons.")
						gobalMap = gobalMap.Merge(unknownScanner, difference)
						knownScanners[unknownScanner.Id()] = difference
						break
					}
					unknownScanner = unknownScanner.RotateZClockwise()

				}

				if overplaps {
					break
				}
				unknownScanner = unknownScanner.RotateYClockwise()
			}

			if overplaps {
				break
			}
			unknownScanner = unknownScanner.RotateXClockwise()
		}

		if !overplaps {
			log.Debug("Scanner does not overlap known region.")
			unknownScanners = append(unknownScanners[1:], unknownScanners[0])
		} else {
			unknownScanners = unknownScanners[1:]
		}
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
