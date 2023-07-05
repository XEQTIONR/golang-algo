package algo

import "fmt"

func InsertionSort(arr []int64, asc bool) ([]int64, int, int) {

	fmt.Println("Insertion Sort.")
	fmt.Printf("Unsorted: %v\n\n", arr)

	inserts := 0
	comparisons := 0
	sorted := []int64{arr[0]}


	for i:= 1; i< len(arr); i++ {
		sorted, comparisons = insert(sorted, arr[i], asc, comparisons)
		inserts++
	}

	fmt.Printf("Sorted: %v\n", sorted)
	return sorted, inserts, comparisons
}

func insert(arr2 []int64, val int64, asc bool, comparisons int) ([]int64, int) {

	temp := []int64{}
	inserted := false
	
	for j, v := range arr2 {
		comparisons++
		if (asc && ( val > v )) || ( !asc && (val < v) ) {
			temp = append(temp, v)
		} else {
			temp = append(temp, val)
			temp = append(temp, arr2[j:]...)
			inserted = true
			break
		}
	}

	if !inserted {
		temp = append(temp, val)
	}

	return temp, comparisons
}