package day06

import log "github.com/sirupsen/logrus"

const birthPeriod = 7

var cache = make(map[int]int)

type lanternfish struct {
	timer int
}

func (fish *lanternfish) Init(timer int) {
	fish.timer = timer
}

func (parent *lanternfish) AncestorsAfter(days int) int {
	ancestors := 1

	if parent.timer < days {
		firstBirthAt := days - parent.timer - 1

		log.Debugf("Checking cache for ancestors of fish having first birth with %v days remaining.", firstBirthAt)
		cachedAncestors, exists := cache[firstBirthAt]

		if exists {
			log.Debugf("Cache hit for key \"%v\".", firstBirthAt)
			ancestors = cachedAncestors
		} else {
			log.Debugf("Cache miss for key \"%v\". Calculating value.", firstBirthAt)

			numBirths := (firstBirthAt / birthPeriod) + 1
			log.Tracef("numBirths = %v", numBirths)

			for i := 0; i < numBirths; i++ {
				child := &lanternfish{}
				child.Init(8)
				remainingDays := firstBirthAt - (i * birthPeriod)
				log.Debugf("Counting ancestors of child %v with %v days remaining.", i, remainingDays)
				ancestors += child.AncestorsAfter(remainingDays)
			}

			cache[firstBirthAt] = ancestors
		}
	}

	log.Tracef("ancestors = %v", ancestors)
	return ancestors
}
