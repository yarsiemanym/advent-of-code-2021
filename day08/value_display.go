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

func (display *valueDisplay) SetSignals(firstDigitSignals string, secondDigitSignals string, thirdDigitSignals string,
	forthDigitSignals string, signalMapper *signalMapper) {

	display.Digits[0].SetSignals(firstDigitSignals, signalMapper)
	display.Digits[1].SetSignals(secondDigitSignals, signalMapper)
	display.Digits[2].SetSignals(thirdDigitSignals, signalMapper)
	display.Digits[3].SetSignals(forthDigitSignals, signalMapper)
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
