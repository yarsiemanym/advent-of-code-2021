package common

import (
	"os"
	"testing"
)

func Test_EnsureInputExists(t *testing.T) {
	InitSession()

	expectedInputFile := "./day01/input.txt"
	os.Remove(expectedInputFile)

	year := 2020
	day := 1
	actualInputFile := EnsureInputExists(year, day)

	if actualInputFile != expectedInputFile {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedInputFile, actualInputFile)
	}

	os.Remove("day01/input.txt")
	os.Remove("day01")
}
