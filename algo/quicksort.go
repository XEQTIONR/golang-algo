package algo

import (
	"example/gosort/types"
	"fmt"
	"reflect"
)

func Quicksort[T types.Number](a []T, asc bool) ([]T, int, int) {
	fmt.Println("Quicksort Sort.")
	fmt.Printf("Unsorted: %v\n\n", a)

	quicksort(a, asc)

	fmt.Printf("Sorted: %v\n", a)

	return a, 0, 0
}

func quicksort[T types.Number](a2 []T, asc bool) {
	if len(a2) > 1 {
		q := partition(a2, asc)
		quicksort(a2[:q], asc)
		quicksort(a2[q+1:], asc)
	}
}

func partition[T types.Number](arr []T, asc bool) (int) {
	x := arr[len(arr) - 1]
	i := -1
	swap := reflect.Swapper(arr)

	for j:=0; j<(len(arr)-1); j++ {
		if ( (asc && arr[j] < x) || (!asc && arr[j] >= x) ) {
			i++
			swap(i, j)
		}
	}
	swap(i+1, len(arr)-1)

	return i+1
}