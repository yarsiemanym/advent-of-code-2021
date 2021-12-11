package day04

import "testing"

func Test_Generator_Init(t *testing.T) {
	numbers := []int{1, 3, 5, 7, 9}
	generator := numberGenerator{}

	generator.Init(numbers)

	if generator.pointer != 0 {
		t.Errorf("Expected 0 but got %v.", generator.pointer)
	}

	for i := 0; i < len(numbers); i++ {
		if generator.numbers[i] != numbers[i] {
			t.Errorf("Expected %v but got %v.", numbers[i], generator.numbers[i])
		}
	}
}

func Test_Generator_Next(t *testing.T) {
	numbers := []int{1, 3, 5, 7, 9}
	generator := numberGenerator{}
	generator.Init(numbers)

	for i := 0; i < len(numbers); i++ {
		next := generator.Next()

		if next != numbers[i] {
			t.Errorf("Expected %v but got %v.", numbers[i], next)
		}
	}
}
