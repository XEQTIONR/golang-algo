package algo

import (
	"fmt"
	"reflect"
)

type Heap struct {
	array []int64
	size int
}

func Heapsort(arr []int64, asc bool) ([]int64, int, int) {
	heap := Heap{ array: arr, size: len(arr) }
	swap := reflect.Swapper(heap.array)

	fmt.Println("Heapsort");
	fmt.Printf("Before sort: %v\n", arr)

	heap.buildHeap(heap.size, asc)
	for i:= len(heap.array) - 1; i>0; i-- {
		swap(0, i)
		heap.size--
		heap.heapify(0, asc)
	}

	fmt.Printf("After sort: %v\n", heap.array)

	return heap.array, 0, 0
}

func parent(i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func(h *Heap) heapify(i int, asc bool) {
	swap := reflect.Swapper(h.array)
	l := left(i)
	r := right(i)
	var extreme int; // largest (asc) / smallest (desc)

	if l < h.size && (asc && h.array[l] >= h.array[i] || !asc && h.array[l] < h.array[i] ) {
		extreme = l
	} else {
		extreme = i
	}

	if r < h.size && (asc && h.array[r] >= h.array[extreme] || !asc && h.array[r] < h.array[extreme]) {
		extreme = r
	}

	if extreme != i {
		swap(i, extreme)
		h.heapify(extreme, asc)
	}
}

func (h *Heap) buildHeap(n int, asc bool) {
	h.size = n
	for i:= n/2; i>=0; i-- {
		h.heapify(i, asc)
	}
}