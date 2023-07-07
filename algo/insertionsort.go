package algo

import "fmt"

func InsertionSort(arr []int64, asc bool) ([]int64, int, int) {

	fmt.Println("Insertion Sort.")
	fmt.Printf("Unsorted: %v\n\n", arr)

	comparisons := 0
	swaps := 0

	for i:=1; i<len(arr); i++ {
		elem := arr[i]
		j:= i-1
		for j>=0 {
			comparisons++
			if (asc && arr[j] > elem) || ( !asc && arr[j] <= elem) { // if we werent counting comparisons this could be in the loop condition.
				arr[j+1] = arr[j]
				swaps++
				j--
			} else {
				break
			}
		}
		arr[j+1] = elem
	}

	fmt.Printf("Sorted: %v\n",  arr)

	return arr, swaps, comparisons

}