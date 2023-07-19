package dynamic_programming

func RecursiveRodCutting(length int, list map[int]int) (int, int) {
	calls := 1 // counting the call in the current instance

	if length == 0 {
		return 0, calls
	}

	max := -1
	for i:=1; i<=length; i++ {
		p, c :=  RecursiveRodCutting(length-i, list)
		part := p + list[i]
		calls += c // add all the calls made by descendants

		if part > max {
			max = part
		}

	}
	return max, calls // the solution and the # of recursive calls required
}

func MemoisedRodCutting(length int, list map[int]int) (int, int) {

	memo := make(map[int]int)

	return memoisedRodCutting(length, list, memo)

}

func memoisedRodCutting(length int, list, memo map[int]int) (int, int) {
	calls := 1 // counting the call in the current instance

	if length == 0 {
		return 0, calls
	}

	val, ok := memo[length]

	if ok {
		return val, calls
	}

	max := -1
	for i:=1; i<=length; i++ {
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