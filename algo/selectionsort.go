package algo

import (
	"fmt"
	"reflect"
)

func SelectionSort(arr []int64, asc bool) ([]int64, int, int) {
	
	fmt.Println("Selection Sort.")
	fmt.Printf("Unsorted: %v\n\n", arr)

	swapper := reflect.Swapper(arr)
	swaps := 0
	comparisons := 0

	for i := range arr[:len(arr) - 1] {
		extreme := arr[i]
		extreme_index := i

		for j:=i+1; j<len(arr); j++ {
			
			compare := arr[j] < extreme
			if (compare && asc) || !(compare || asc)  {
				extreme = arr[j]
				extreme_index = j
			}
			comparisons++
		}

		swapper(i, extreme_index)
		swaps++
	}

	fmt.Printf("\nAfter sort: %v\n", arr)

	return arr, swaps, comparisons
}