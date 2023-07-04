package main

import (
	"fmt"
	"gosort/algo"
)

func main() {
	arr := []int64{11,1, 16, 6, 4, 5, 3, 15, 10}
	sorted, _, _ := algo.BubbleSort(arr, true)
	fmt.Println(sorted)
}