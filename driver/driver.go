package main

import (
	"example/algo/sorting"
	"fmt"
)

func main() {
	arr := []int64{11,1, 16, 6, 4, 5, 3, 15, 10}
	
	_, swaps, comparisons := sorting.Heapsort(arr, true)

	fmt.Printf("Swaps/Inserts: %v, Comparisons: %v", swaps, comparisons)
}