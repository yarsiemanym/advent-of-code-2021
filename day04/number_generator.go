package day04

type numberGenerator struct {
	numbers []int
	pointer int
}

func (generator *numberGenerator) Init(numbers []int) {
	generator.numbers = numbers
	generator.pointer = 0
}

func (generator *numberGenerator) Next() int {
	number := generator.numbers[generator.pointer]
	generator.pointer++
	return number
}
