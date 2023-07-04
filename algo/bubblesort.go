package algo

import (
	"fmt"
	"reflect"
)

func BubbleSort(arr []int64, asc bool) ([]int64, int, int) {

	fmt.Println("Bubble Sort.")
	fmt.Printf("Unsorted: %v\n\n", arr)

	swapper := reflect.Swapper(arr)
	swaps := 0
	comparisons := 0

	for i := range arr[:len(arr)-1] {

		for j := range arr[:len(arr)-i-1] {

			compare := arr[j] > arr[j+1]

			if (compare && asc) || !(compare || asc) {
				swapper(j, j+1)
				swaps++
			}
			comparisons++
		}
		fmt.Printf("After pass %v : %v\n", i+1, arr)
	}
	return arr, swaps, comparisons
}