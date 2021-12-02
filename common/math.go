package common

func SumInt(values ...int) int {
	return SumIntSlice(values)
}

func SumIntSlice(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}
