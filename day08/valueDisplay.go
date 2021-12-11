package day08

import (
	"fmt"
)

type valueDisplay struct {
	digits []*digitDisplay
}

func NewValueDisplay() *valueDisplay {
	return &valueDisplay{
		digits: []*digitDisplay{
			NewDigitDisplay(0),
			NewDigitDisplay(1),
			NewDigitDisplay(2),
			NewDigitDisplay(3),
		},
	}
}

func (display *valueDisplay) SetSignals(firstDigitSignals string, secondDigitSignals string, thirdDigitSignals string,
	forthDigitSignals string, signalMapper *signalMapper) {

	display.digits[0].SetSignals(firstDigitSignals, signalMapper)
	display.digits[1].SetSignals(secondDigitSignals, signalMapper)
	display.digits[2].SetSignals(thirdDigitSignals, signalMapper)
	display.digits[3].SetSignals(forthDigitSignals, signalMapper)
}

func (display *valueDisplay) Render() string {
	ouput := ""

	for _, digit := range display.digits {
		ouput += fmt.Sprintf("  %v:   ", digit.id)
	}
	ouput += "\n"

	for i := 0; i < 7; i++ {
		line := ""

		for _, digit := range display.digits {
			line += digit.Render(i) + " "
		}
		ouput += line + "\n"
	}

	return ouput
}

func (display *valueDisplay) NumericValue() *int {
	numericValue := new(int)

	for _, digit := range display.digits {
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
