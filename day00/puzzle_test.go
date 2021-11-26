package day00

import (
	"strconv"
	"testing"
	"time"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func Test_Solve_Input1(t *testing.T) {
	input := "test1.txt"

	expectedName := "Joe Schmoe"
	birthday, err := time.Parse(common.ShortDateFormat, "1983-11-24")
	common.Check(err)
	expectedAge := int(time.Now().Sub(birthday).Hours() / 24 / 365)

	actualName, actualAge := Solve(input)

	if actualName != expectedName {
		t.Errorf("Expected:\n%v\nActual:\n%v", expectedName, actualName)
	}

	if actualAge != strconv.Itoa(expectedAge) {
		t.Errorf("Expected:\n%v\nActual:\n%v", expectedAge, actualAge)
	}
}
