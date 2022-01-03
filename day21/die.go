package day21

type Die interface {
	RollN(n int) int
	Roll() int
	NumberOfRolls() int
}
