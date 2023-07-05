package main

import (
	"gosort/algo"
	"fmt"
)

func main() {
	arr := []int64{11,1, 16, 6, 4, 5, 3, 15, 10}
	
	_, swaps, comparisons := algo.InsertionSort(arr, true)

	fmt.Printf("Swaps/Inserts: %v, Comparisons: %v", swaps, comparisons)
}