package algo

import (
	"fmt"
	"reflect"
)

func BubbleSort(array []int64, asc bool) ([]int64, int, int) {

	fmt.Println("Bubble Sort.")

	swapper := reflect.Swapper(array)
	swaps := 0
	loops := 0

	for i := range array[:len(array)-1] {

		for j := range array[:len(array)-i-1] {

			compare := array[j] > array[j+1]

			if (compare && asc) || !(compare || asc) {
				swapper(j, j+1)
				swaps++
			}
			loops++
		}
		fmt.Printf("After pass %v : %v\n", i, array)
	}
	return array, swaps, loops
}