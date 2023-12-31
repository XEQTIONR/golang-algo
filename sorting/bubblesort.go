package sorting

import (
	"fmt"
	"reflect"
	"example/algo/types"
)

func BubbleSort[T types.Number](arr []T, asc bool) ([]T, int, int) {

	fmt.Println("Bubble Sort.")
	fmt.Printf("Unsorted: %v\n\n", arr)

	swapper := reflect.Swapper(arr)
	swaps := 0
	comparisons := 0

	for i := range arr[:len(arr)-1] {

		swapped := false
		for j := range arr[:len(arr)-i-1] {

			compare := arr[j] > arr[j+1]

			if (compare && asc) || !(compare || asc) {
				swapper(j, j+1)
				swaps++
				swapped = true
			}
			comparisons++
		}
		fmt.Printf("After pass %v : %v\n", i+1, arr)
		if !swapped {
			break
		}
	}
	return arr, swaps, comparisons
}