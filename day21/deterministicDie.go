package day21

import log "github.com/sirupsen/logrus"

type DeterminisiticDie struct {
	max   int
	last  int
	rolls int
}

func NewDeterministicDie(max int) *DeterminisiticDie {
	return &DeterminisiticDie{
		max:   max,
		last:  0,
		rolls: 0,
	}
}

func (die *DeterminisiticDie) RollN(n int) int {
	log.Debugf("Rolling %d die.", n)

	sum := 0

	for i := 0; i < n; i++ {
		sum += die.Roll()
	}

	log.Debugf("Sum of rolls is %d.", sum)

	return sum
}

func (die *DeterminisiticDie) Roll() int {
	log.Trace("Rolling die.")
	roll := die.last + 1

	if roll > die.max {
		roll = roll - die.max
	}

	die.rolls++
	die.last = roll

	log.Tracef("Rolled %d.", roll)

	return roll
}

func (die *DeterminisiticDie) NumberOfRolls() int {
	return die.rolls
}
