package day19

import (
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type Scanner struct {
	id      int
	beacons []*common.Point
}

func NewScanner(id int, beacons []*common.Point) *Scanner {
	return &Scanner{
		id:      id,
		beacons: beacons,
	}
}

func (scanner *Scanner) Id() int {
	return scanner.id
}

func (scanner *Scanner) Beacons() []*common.Point {
	return scanner.beacons
}

func (scanner *Scanner) RotateX() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateX()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) RotateY() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateY()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) RotateZ() *Scanner {
	rotatedPoints := make([]*common.Point, len(scanner.Beacons()))

	for i, point := range scanner.Beacons() {
		rotatedPoints[i] = point.RotateZ()
	}

	return NewScanner(scanner.Id(), rotatedPoints)
}

func (scanner *Scanner) DetectBeaconOverlap(other *Scanner) (bool, *common.Point) {
	for _, anchor1 := range scanner.Beacons() {
		for _, anchor2 := range other.Beacons() {
			referenceDifference := anchor1.Difference(anchor2)
			overlappingBeacons := []*common.Point{}

			for _, beacon1 := range scanner.Beacons() {
				for _, beacon2 := range other.Beacons() {
					difference := beacon1.Difference(beacon2)
					if *difference == *referenceDifference {
						overlappingBeacons = append(overlappingBeacons, beacon2)
					}
				}
			}

			if len(overlappingBeacons) >= 12 {
				return true, referenceDifference
			}
		}
	}

	return false, nil
}

func (scanner *Scanner) Merge(other *Scanner, difference *common.Point) *Scanner {
	newBeacons := make([]*common.Point, len(scanner.Beacons()))
	copy(newBeacons, scanner.Beacons())

	for _, beacon1 := range other.Beacons() {
		newBeacon := beacon1.Move(difference)
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
