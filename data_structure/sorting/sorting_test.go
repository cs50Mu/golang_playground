package sorting

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{7, 2, 6, 3, 9, 5}
	fmt.Printf("before msort: %+v\n", a)
	mergeSort(a)
	fmt.Printf("after msort: %+v\n", a)
}

func TestQuickSort(t *testing.T) {
	a := []int{7, 2, 6, 3, 9, 5}
	fmt.Printf("before qsort: %+v\n", a)
	quickSort(a)
	fmt.Printf("after qsort: %+v\n", a)
}

func TestHeapSort(t *testing.T) {
	a := []int{7, 2, 6, 3, 9, 5}
	fmt.Printf("before hsort: %+v\n", a)
	heapSort(a)
	fmt.Printf("after hsort: %+v\n", a)
}
