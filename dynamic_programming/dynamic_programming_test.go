package dynamic_programming

import (
	"reflect"
	"testing"
	"fmt"
)

func TestRecursiveRodCutting(t *testing.T) {

	rMax4, _ := RecursiveRodCutting(4, priceList(4))
	rMax7, _ := RecursiveRodCutting(7, priceList(7))
	rMax10, _ := RecursiveRodCutting(10, priceList(10))
	mMax4, _ := MemoisedRodCutting(4, priceList(4))
	mMax7, _ := MemoisedRodCutting(7, priceList(7))
	mMax10, _ := MemoisedRodCutting(10, priceList(10))

	if !reflect.DeepEqual(rMax4, 10) {
		t.Errorf(errorMessage(10, rMax4))
	}

	if !reflect.DeepEqual(rMax7, 18) {
		t.Errorf(errorMessage(18, rMax7))
	}

	if !reflect.DeepEqual(rMax10, 30) {
		t.Errorf(errorMessage(18, rMax10))
	}

	if !reflect.DeepEqual(mMax4, 10) {
		t.Errorf(errorMessage(10, mMax4))
	}

	if !reflect.DeepEqual(mMax7, 18) {
		t.Errorf(errorMessage(18, mMax7))
	}

	if !reflect.DeepEqual(mMax10, 30) {
		t.Errorf(errorMessage(18, mMax10))
	}

}

func priceList(n int) map[int]int {
	var list map[int]int
	switch n {
		case 4:
			list = map[int]int{
				1: 1,
				2: 5,
				3: 8,
				4: 9,
			}
		
		case 7:
			list = map[int]int{
				1: 1,
				2: 5,
				3: 8,
				4: 9,
				5: 10,
				6: 17,
				7: 17,
			}

		case 10:
			list = map[int]int{
				1: 1,
				2: 5,
				3: 8,
				4: 9,
				5: 10,
				6: 17,
				7: 17,
				8: 20,
				9: 24,
				10: 30,
			}
	}

	return list
}

func errorMessage(actual, expected int) string {
	return fmt.Sprintf("Expected: %v Actual: %v\n", expected, actual)
}