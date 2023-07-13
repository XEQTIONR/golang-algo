package algo

import (
	"example/gosort/types"
	"fmt"
)

func MergeSort[T types.Number](arr []T, asc bool) ([]T, int, int) {

	fmt.Println("Merge Sort.")
	fmt.Printf("Unsorted: %v\n\n", arr)

	sorted := sort(arr, asc)
	fmt.Printf("Sorted: %v\n", sorted)

	return sorted, 0, 0
}

func sort[T types.Number](arr []T, asc bool) []T {
	switch len(arr) {
		case 1:
			return arr
		case 2:
			if (asc && arr[0] < arr[1]) || (!asc && arr[0] >= arr[1]) {
				return arr
			} else {
				return []T{arr[1], arr[0]}
			}
		default:
			left :=  sort(arr[ :(len(arr)/2)], asc)
			right := sort(arr[len(arr)/2: ], asc)
			return merge(left, right, asc)
	}
}

func merge[T types.Number](a, b []T, asc bool) []T {
	out:= make([]T, len(a) + len(b))
	ai, bi := 0, 0
	
	for i, _ := range out {
		if ai >= len(a) {
			out[i] = b[bi]
			bi++
		} else if bi >= len(b) {
			out[i] = a[ai]
			ai++
		} else if (asc && a[ai] < b[bi]) || (!asc && a[ai] >= b[bi]) {
			out[i] = a[ai]
			ai++
		} else {
			out[i] = b[bi]
			bi++
		}
	}
	return out

}