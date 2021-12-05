package common

import "math"

func SumIntVariadic(values ...int) int {
	return SumIntSlice(values)
}

func SumIntSlice(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}

func MaxIntVariadic(values ...int) int {
	return MaxIntSlice(values)
}

func MaxIntSlice(values []int) int {
	maxValue := math.MinInt

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func MinIntVariadic(values ...int) int {
	return MinIntSlice(values)
}

func MinIntSlice(values []int) int {
	minValue := math.MaxInt

	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func Reduce(numerator int, denominator int) (int, int) {
	if numerator == 0 {
		return 0, int(math.Copysign(1, float64(denominator)))
	}

	if denominator == 0 {
		return int(math.Copysign(1, float64(numerator))), 0
	}

	gcd := GreatestCommonDenominator(numerator, denominator)
	return numerator / gcd, denominator / gcd
}

func GreatestCommonDenominator(a int, b int) int {

	if a == 0 || b == 0 {
		return 1
	}

	if a < 0 {
		a = int(math.Abs(float64(a)))
	}

	if b < 0 {
		b = int(math.Abs(float64(b)))
	}

	if a > b {
		a, b = b, a
	}

	if b%a == 0 {
		return a
	} else if a > b {
		return GreatestCommonDenominator(a-b, b)
	} else {
		return GreatestCommonDenominator(a, b-a)
	}
}
