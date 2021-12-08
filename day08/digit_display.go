package day08

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type digitDisplay struct {
	id       int
	segments map[rune]*wireSegment
}

func NewDigitDisplay(id int) *digitDisplay {
	return &digitDisplay{
		id: id,
		segments: map[rune]*wireSegment{
			'a': {
				signal: 'a',
				isOn:   false,
			},
			'b': {
				signal: 'b',
				isOn:   false,
			},
			'c': {
				signal: 'c',
				isOn:   false,
			},
			'd': {
				signal: 'd',
				isOn:   false,
			},
			'e': {
				signal: 'e',
				isOn:   false,
			},
			'f': {
				signal: 'f',
				isOn:   false,
			},
			'g': {
				signal: 'g',
				isOn:   false,
			},
		},
	}
}

func (display *digitDisplay) Clear() {
	for _, segment := range display.segments {
		segment.isOn = false
	}
}

func (display *digitDisplay) SetSignals(signals string) {
	for _, signal := range signals {
		for _, segment := range display.segments {
			if segment.signal == signal {
				segment.isOn = true
			}
		}
	}
}

func (display *digitDisplay) Render(lineNumber int) string {
	line := ""

	switch lineNumber {
	case 0:
		a := '.'
		if display.segments['a'].isOn {
			a = display.segments['a'].signal
		}

		line = fmt.Sprintf(" %c%c%c%c ", a, a, a, a)
	case 1, 2:
		b := '.'
		if display.segments['b'].isOn {
			b = display.segments['b'].signal
		}

		c := '.'
		if display.segments['c'].isOn {
			c = display.segments['c'].signal
		}

		line = fmt.Sprintf("%c    %c", b, c)
	case 3:
		d := '.'
		if display.segments['d'].isOn {
			d = display.segments['d'].signal
		}

		line = fmt.Sprintf(" %c%c%c%c ", d, d, d, d)
	case 4, 5:
		e := '.'
		if display.segments['e'].isOn {
			e = display.segments['e'].signal
		}

		f := '.'
		if display.segments['f'].isOn {
			f = display.segments['f'].signal
		}

		line = fmt.Sprintf("%c    %c", e, f)
	case 6:
		g := '.'
		if display.segments['g'].isOn {
			g = display.segments['g'].signal
		}

		line = fmt.Sprintf(" %c%c%c%c ", g, g, g, g)
	default:
		log.Fatalf("%v is not a valid line number.", lineNumber)
	}

	return line
}

func (display *digitDisplay) NumericValue() *int {
	numericValue := new(int)

	if display.IsZero() {
		*numericValue = 0
	} else if display.IsOne() {
		*numericValue = 1
	} else if display.IsTwo() {
		*numericValue = 2
	} else if display.IsThree() {
		*numericValue = 3
	} else if display.IsFour() {
		*numericValue = 4
	} else if display.IsFive() {
		*numericValue = 5
	} else if display.IsSix() {
		*numericValue = 6
	} else if display.IsSeven() {
		*numericValue = 7
	} else if display.IsEight() {
		*numericValue = 8
	} else if display.IsNine() {
		*numericValue = 9
	} else {
		numericValue = nil
	}

	return numericValue
}

func (display *digitDisplay) IsZero() bool {
	return display.segments['a'].isOn &&
		display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		!display.segments['d'].isOn &&
		display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		display.segments['g'].isOn
}

func (display *digitDisplay) IsOne() bool {
	return !display.segments['a'].isOn &&
		!display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		!display.segments['d'].isOn &&
		!display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		!display.segments['g'].isOn
}

func (display *digitDisplay) IsTwo() bool {
	return display.segments['a'].isOn &&
		!display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		display.segments['d'].isOn &&
		display.segments['e'].isOn &&
		!display.segments['f'].isOn &&
		display.segments['g'].isOn
}

func (display *digitDisplay) IsThree() bool {
	return display.segments['a'].isOn &&
		!display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		display.segments['d'].isOn &&
		!display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		display.segments['g'].isOn
}

func (display *digitDisplay) IsFour() bool {
	return !display.segments['a'].isOn &&
		display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		display.segments['d'].isOn &&
		!display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		!display.segments['g'].isOn
}

func (display *digitDisplay) IsFive() bool {
	return display.segments['a'].isOn &&
		display.segments['b'].isOn &&
		!display.segments['c'].isOn &&
		display.segments['d'].isOn &&
		!display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		display.segments['g'].isOn
}

func (display *digitDisplay) IsSix() bool {
	return display.segments['a'].isOn &&
		display.segments['b'].isOn &&
		!display.segments['c'].isOn &&
		display.segments['d'].isOn &&
		display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		display.segments['g'].isOn
}

func (display *digitDisplay) IsSeven() bool {
	return display.segments['a'].isOn &&
		!display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		!display.segments['d'].isOn &&
		!display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		!display.segments['g'].isOn
}

func (display *digitDisplay) IsEight() bool {
	return display.segments['a'].isOn &&
		display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		display.segments['d'].isOn &&
		display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		display.segments['g'].isOn
}

func (display *digitDisplay) IsNine() bool {
	return display.segments['a'].isOn &&
		display.segments['b'].isOn &&
		display.segments['c'].isOn &&
		display.segments['d'].isOn &&
		!display.segments['e'].isOn &&
		display.segments['f'].isOn &&
		display.segments['g'].isOn
}
