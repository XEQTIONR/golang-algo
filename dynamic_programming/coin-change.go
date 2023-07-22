package dynamic_programming

import "math"

func RecursiveCoinChange(total int, denom []int) (int, int) {

	min := math.MaxInt
	calls := 1

	if total == 0 {
		return 0, calls
	}

	for _, d := range denom {
		if d <= total {
				m, c := RecursiveCoinChange(total-d, denom)
				m++ // add one coin of d denomination
				calls+= c

				if m < min {
					min = m
				}
		}
	}
	return min, calls
}

func MemoisedCoinChange(total int, denom []int) (int, int) {

	memo := make(map[int]int)
	return memoisedCoinChange(total, denom, memo)
}

func memoisedCoinChange(total int, denom []int, memo map[int]int) (int, int) {

	min := math.MaxInt
	calls := 1
	var m int

	if total == 0 {
		return 0, calls
	}

	for _, d := range denom {
		if d <= total {
			val, ok  := memo[total-d]

			if ok {
				m = val
			} else {
				n, c := memoisedCoinChange(total-d, denom, memo)
				m = n + 1 // add one coin of d denomination
				calls+= c
			}

			if m < min {
				min = m
				memo[total-d] = min
			}
		}
	}
	return min, calls

}