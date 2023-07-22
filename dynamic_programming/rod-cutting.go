package dynamic_programming

import "example/algo/types"

func RecursiveRodCutting[L types.Number, T types.Number](length L, list map[L]T) (T, int) {
	calls := 1 // counting the call in the current instance

	if length == L(0) {
		return T(0), calls
	}

	max := T(-1)
	for i:=L(1); i<=length; i++ {
		p, c :=  RecursiveRodCutting(length-i, list)
		part := p + list[i]
		calls += c // add all the calls made by descendants

		if part > max {
			max = part
		}

	}
	return max, calls // the solution and the # of recursive calls required
}

func MemoisedRodCutting[L types.Number, T types.Number](length L, list map[L]T) (T, int) {

	memo := make(map[L]T)

	return memoisedRodCutting(length, list, memo)

}

func memoisedRodCutting[L types.Number, T types.Number](length L, list, memo map[L]T) (T, int) {
	calls := 1 // counting the call in the current instance

	if length == L(0) {
		return T(0), calls
	}

	val, ok := memo[length]

	if ok {
		return val, calls
	}

	max := T(-1)
	for i:= L(1); i<=length; i++ {
		p, c :=  memoisedRodCutting(length-i, list, memo)
		part := p + list[i]
		calls += c // add all the calls made by descendants

		if part > max {
			max = part
		}

	}

	memo[length] = max

	return max, calls

}