package algo

import "fmt"

func MergeSort(arr []int64, asc bool) ([]int64, int, int) {

	fmt.Println("Merge Sort.")
	fmt.Printf("Unsorted: %v\n\n", arr)

	sorted := sort(arr, asc)
	fmt.Printf("Sorted: %v\n", sorted)

	return sorted, 0, 0
}

func sort(arr []int64, asc bool) []int64{
	switch len(arr) {
		case 1:
			return arr
		case 2:
			if (asc && arr[0] < arr[1]) || (!asc && arr[0] >= arr[1]) {
				return arr
			} else {
				return []int64{arr[1], arr[0]}
			}
		default:
			left, right := split(arr)
			left = sort(left, asc)
			right = sort(right, asc)
			return merge(left, right, asc)
	}
}

func split(arr []int64) ([]int64, []int64) {
	return arr[:(len(arr)/2)], arr[len(arr)/2:]
}

func merge(a, b []int64, asc bool) []int64 {
	out:= make([]int64, len(a) + len(b))
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