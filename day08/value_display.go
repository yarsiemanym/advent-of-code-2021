package day08

import (
	"fmt"
)

type valueDisplay struct {
	Digits []*digitDisplay
}

func NewValueDisplay() *valueDisplay {
	return &valueDisplay{
		Digits: []*digitDisplay{
			NewDigitDisplay(0),
			NewDigitDisplay(1),
			NewDigitDisplay(2),
			NewDigitDisplay(3),
		},
	}
}

func (display *valueDisplay) SetSignals(first string, second string, third string, forth string, signalMapper *signalMapper) {
	if signalMapper != nil {
		first = signalMapper.MapSignalPattern(first)
		second = signalMapper.MapSignalPattern(second)
		third = signalMapper.MapSignalPattern(third)
		forth = signalMapper.MapSignalPattern(forth)
	}

	display.Digits[0].SetSignals(first)
	display.Digits[1].SetSignals(second)
	display.Digits[2].SetSignals(third)
	display.Digits[3].SetSignals(forth)
}

func (display *valueDisplay) Render() string {
	ouput := ""

	for _, digit := range display.Digits {
		ouput += fmt.Sprintf("  %v:   ", digit.id)
	}
	ouput += "\n"

	for i := 0; i < 7; i++ {
		line := ""

		for _, digit := range display.Digits {
			line += digit.Render(i) + " "
		}
		ouput += line + "\n"
	}

	return ouput
}

func (display *valueDisplay) NumericValue() *int {
	numericValue := new(int)

	for _, digit := range display.Digits {
		digitNumericValue := digit.NumericValue()

		if digitNumericValue == nil {
			numericValue = nil
			break
		} else {
			*numericValue = (10 * *numericValue) + *digitNumericValue
		}
	}

	return numericValue
}
