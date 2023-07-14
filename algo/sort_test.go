package algo

import (
	"example/gosort/types"
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	sortedInts, _, _ := BubbleSort(intArray(), true)
	sortedInt64s, _, _ := BubbleSort(int64Array(), true)
	sortedFloat64s, _, _ := BubbleSort(float64Array(), true)

	if !reflect.DeepEqual(sortedInts, sortedIntArray()) {
		t.Errorf(errorMessage(sortedInts, sortedIntArray()))
	}

	if !reflect.DeepEqual(sortedInt64s, sortedInt64Array()) {
		t.Errorf(errorMessage(sortedInt64s, sortedInt64Array()))
	}

	if !reflect.DeepEqual(sortedFloat64s, sortedFloat64Array()) {
		t.Errorf(errorMessage(sortedFloat64s, sortedFloat64Array()))
	}
}

func TestInsertionSort(t *testing.T) {

	sortedInts, _, _ := InsertionSort(intArray(), true)
	sortedInt64s, _, _ := InsertionSort(int64Array(), true)
	sortedFloat64s, _, _ := InsertionSort(float64Array(), true)

	if !reflect.DeepEqual(sortedInts, sortedIntArray()) {
		t.Errorf(errorMessage(sortedInts, sortedIntArray()))
	}

	if !reflect.DeepEqual(sortedInt64s, sortedInt64Array()) {
		t.Errorf(errorMessage(sortedInt64s, sortedInt64Array()))
	}

	if !reflect.DeepEqual(sortedFloat64s, sortedFloat64Array()) {
		t.Errorf(errorMessage(sortedFloat64s, sortedFloat64Array()))
	}
}

func TestSelectionSort(t *testing.T) {

	sortedInts, _, _ := SelectionSort(intArray(), true)
	sortedInt64s, _, _ := SelectionSort(int64Array(), true)
	sortedFloat64s, _, _ := SelectionSort(float64Array(), true)

	if !reflect.DeepEqual(sortedInts, sortedIntArray()) {
		t.Errorf(errorMessage(sortedInts, sortedIntArray()))
	}

	if !reflect.DeepEqual(sortedInt64s, sortedInt64Array()) {
		t.Errorf(errorMessage(sortedInt64s, sortedInt64Array()))
	}

	if !reflect.DeepEqual(sortedFloat64s, sortedFloat64Array()) {
		t.Errorf(errorMessage(sortedFloat64s, sortedFloat64Array()))
	}
}

func TestMergeSort(t *testing.T) {

	sortedInts, _, _ := MergeSort(intArray(), true)
	sortedInt64s, _, _ := MergeSort(int64Array(), true)
	sortedFloat64s, _, _ := MergeSort(float64Array(), true)

	if !reflect.DeepEqual(sortedInts, sortedIntArray()) {
		t.Errorf(errorMessage(sortedInts, sortedIntArray()))
	}

	if !reflect.DeepEqual(sortedInt64s, sortedInt64Array()) {
		t.Errorf(errorMessage(sortedInt64s, sortedInt64Array()))
	}

	if !reflect.DeepEqual(sortedFloat64s, sortedFloat64Array()) {
		t.Errorf(errorMessage(sortedFloat64s, sortedFloat64Array()))
	}
}

func TestHeapsort(t *testing.T) {

	sortedInts, _, _ := Heapsort(intArray(), true)
	sortedInt64s, _, _ := Heapsort(int64Array(), true)
	sortedFloat64s, _, _ := Heapsort(float64Array(), true)

	if !reflect.DeepEqual(sortedInts, sortedIntArray()) {
		t.Errorf(errorMessage(sortedInts, sortedIntArray()))
	}

	if !reflect.DeepEqual(sortedInt64s, sortedInt64Array()) {
		t.Errorf(errorMessage(sortedInt64s, sortedInt64Array()))
	}

	if !reflect.DeepEqual(sortedFloat64s, sortedFloat64Array()) {
		t.Errorf(errorMessage(sortedFloat64s, sortedFloat64Array()))
	}
}

func TestQuicksort(t *testing.T) {

	sortedInts, _, _ := Quicksort(intArray(), true)
	sortedInt64s, _, _ := Quicksort(int64Array(), true)
	sortedFloat64s, _, _ := Quicksort(float64Array(), true)

	if !reflect.DeepEqual(sortedInts, sortedIntArray()) {
		t.Errorf(errorMessage(sortedInts, sortedIntArray()))
	}

	if !reflect.DeepEqual(sortedInt64s, sortedInt64Array()) {
		t.Errorf(errorMessage(sortedInt64s, sortedInt64Array()))
	}

	if !reflect.DeepEqual(sortedFloat64s, sortedFloat64Array()) {
		t.Errorf(errorMessage(sortedFloat64s, sortedFloat64Array()))
	}
}

func intArray() []int {
	return []int{9, 1, 5, 2, 7, 8, 3, 0, 6, 4}
}

func int64Array() []int64 {
	return []int64{9, 1, 5, 2, 7, 8, 3, 0, 6, 4}
}

func float64Array() []float64 {
	return []float64{9.0, 1, 5, 2, 7.8, 8, 3, 0, 6, 4.2}
}

func sortedIntArray() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func sortedInt64Array() []int64 {
	return []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func sortedFloat64Array() []float64 {
	return []float64{0, 1, 2, 3, 4.2, 5, 6, 7.8, 8, 9.0}
}

func errorMessage[T types.Number] (actual, expected []T ) string {
	return fmt.Sprintf("Expected: %v\nActual: %v\n", expected, actual)
}