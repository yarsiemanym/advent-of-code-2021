package day19

import (
	"github.com/Workiva/go-datastructures/set"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type Scanner struct {
	id          int
	fingerprint map[float64][]*common.Point
	beacons     []*common.Point
}

func NewScanner(id int, beacons []*common.Point) *Scanner {
	return &Scanner{
		id:          id,
		fingerprint: fingerprint(beacons),
		beacons:     beacons,
	}
}

func fingerprint(beacons []*common.Point) map[float64][]*common.Point {
	distances := map[float64][]*common.Point{}

	for i := 0; i < len(beacons); i++ {
		beacon1 := beacons[i]
		for j := i + 1; j < len(beacons); j++ {
			beacon2 := beacons[j]
			distance := beacon1.Distance(beacon2)
			distances[distance] = []*common.Point{beacon1, beacon2}
		}
	}
	return distances
}

func (scanner *Scanner) Id() int {
	return scanner.id
}

func (scanner *Scanner) Beacons() []*common.Point {
	return scanner.beacons
}

func (scanner *Scanner) RotateXClockwise() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateXClockwise()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) RotateXCounterClockwise() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateXCounterClockwise()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) RotateYClockwise() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateYClockwise()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) RotateYCounterClockwise() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateYCounterClockwise()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) RotateZClockwise() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateZClockwise()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) RotateZCounterClockwise() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateZCounterClockwise()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) DetectOverlap(other *Scanner) bool {
	sharedPoints := set.New()

	for distance := range scanner.fingerprint {
		points, exists := other.fingerprint[distance]

		if exists {
			sharedPoints.Add(*points[0])
			sharedPoints.Add(*points[1])
		}
	}

	return sharedPoints.Len() >= 12
}

func (scanner *Scanner) Align(other *Scanner) (bool, *Scanner, *common.Point) {
	originalAlignment := other
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				aligned, difference := scanner.IsAligned(other)

				if aligned {
					return aligned, other, difference
				}

				other = other.RotateZClockwise()
			}

			other = other.RotateYClockwise()
		}

		other = other.RotateXClockwise()
	}

	return false, originalAlignment, nil
}

func (scanner *Scanner) IsAligned(other *Scanner) (bool, *common.Point) {
	differences := map[common.Point]int{}

	for _, beacon1 := range scanner.Beacons() {
		for _, beacon2 := range other.Beacons() {
			difference := beacon2.Subtract(beacon1)
			differences[*difference] = differences[*difference] + 1
		}
	}

	for difference, count := range differences {
		if count >= 12 {
			return true, &difference
		}
	}

	return false, nil

}

func (scanner *Scanner) Merge(other *Scanner, difference *common.Point) *Scanner {
	newBeacons := make([]*common.Point, len(scanner.Beacons()))
	copy(newBeacons, scanner.Beacons())

	for _, beacon1 := range other.Beacons() {
		newBeacon := beacon1.Subtract(difference)
		duplicate := false
		for _, beacon2 := range scanner.Beacons() {
			if *newBeacon == *beacon2 {
				duplicate = true
				break
			}
		}

		if !duplicate {
			newBeacons = append(newBeacons, newBeacon)
		}
	}

	return NewScanner(scanner.Id(), newBeacons)
}
