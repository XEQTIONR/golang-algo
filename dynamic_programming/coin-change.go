package dynamic_programming

import (
	"math"
	"example/algo/types"
)

func RecursiveCoinChange[T types.Number](total T, denom []T) (int, int) {
	min := math.MaxInt
	calls := 1

	if total == 0 {
		return 0, calls
	}

	for _, d := range denom {
		if d <= total {
				count, k := RecursiveCoinChange(total-d, denom)
				count++ // add one coin of d denomination
				calls += k

				if count < min {
					min = count
				}
		}
	}
	return min, calls
}

func MemoisedCoinChange[T types.Number](total T, denom []T) (int, int) {
	memo := make(map[T]int)
	return memoisedCoinChange(total, denom, memo)
}

func memoisedCoinChange[T types.Number](total T, denom []T, memo map[T]int) (int, int) {
	min := math.MaxInt
	calls := 1
	var count int

	if total == 0 {
		return 0, calls
	}

	for _, d := range denom {
		if d <= total {
			val, ok  := memo[total-d]

			if ok {
				count = val
			} else {
				n, k := memoisedCoinChange(total-d, denom, memo)
				count = n + 1 // add one coin of d denomination
				calls += k
			}

			if count < min {
				min = count
				memo[total-d] = min
			}
		}
	}
	return min, calls

}