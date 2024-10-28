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

	// b := []int{6, 2, 3, 3, 4, 6, 9, 3, 1}
	b := []int{6, 3, 2, 3, 9, 4, 3, 5}
	fmt.Printf("before MergeSortRec: %+v\n", b)
	ms := NewMergeSort(b)
	ms.Sort()
	fmt.Printf("after MergeSortRec: %+v\n", b)
}

func TestQuickSort(t *testing.T) {
	a := []int{7, 2, 6, 3, 9, 5}
	fmt.Printf("before qsort: %+v\n", a)
	quickSort(a)
	fmt.Printf("after qsort: %+v\n", a)

	b := []int{7, 3, 5, 3, 2, 3}
	fmt.Printf("before QuickSort: %+v\n", b)
	qs := NewQuickSort(b)
	qs.Sort()
	fmt.Printf("after QuickSort: %+v\n", b)
}

func TestHeapSort(t *testing.T) {
	a := []int{7, 2, 6, 3, 9, 5}
	fmt.Printf("before hsort: %+v\n", a)
	heapSort(a)
	fmt.Printf("after hsort: %+v\n", a)

	b := []int{6, 8, 3, 5, 7, 5, 2, 9}
	// b := []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1}
	// b := []int{0, 7, 4, 9, 0, -1, -7, -5}
	fmt.Printf("before HeapSort: %+v\n", b)
	hs := NewHeapSort(b)
	hs.Sort()
	fmt.Printf("after HeapSort: %+v\n", b)
}

func TestHeapPushPop(t *testing.T) {
	// 传入 i < j 为小顶堆
	// 传入 i > j 为大顶堆
	h := NewHeap[int](func(i, j int) bool {
		return i > j
	})

	var got, want int
	h.Push(3)
	h.Push(5)

	got = h.Pop()
	want = 5
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
	fmt.Printf("h.arr: %+v\n", h.arr[:h.size])

	h.Push(50)
	fmt.Printf("h.arr: %+v\n", h.arr[:h.size])
	got = h.Pop()
	fmt.Printf("h.arr: %+v\n", h.arr[:h.size])
	want = 50
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}

	got = h.Pop()
	fmt.Printf("h.arr: %+v\n", h.arr[:h.size])
	want = 3
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestHalveArray(t *testing.T) {
	var want, got int
	var input []int
	input = []int{5, 19, 8, 1}
	want = 3
	got = halveArray(input)
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}

	input = []int{3, 8, 20}
	want = 3
	got = halveArray(input)
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
