package sorting

import (
	"fmt"
	"math/rand"
)

func mergeSort(a []int) {
	aux := make([]int, len(a))
	sort(a, aux, 0, len(a)-1)
}

func merge(a, aux []int, lo, mid, hi int) {
	// copy to aux
	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}

	// merge
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		// 一半已经遍历完的情况
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[i] > aux[j] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func sort(a, aux []int, lo, hi int) {
	// terminate cond
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, aux, lo, mid)
	sort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

////////////////////////////////////////////

func quickSort(a []int) {
	// shuffle
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	quickS(a, 0, len(a)-1)
}

func quickS(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	// partition
	j := partition(a, lo, hi)
	fmt.Printf("j: %v, a: %+v\n", j, a)
	// sort left
	quickS(a, lo, j-1)
	// sort right
	quickS(a, j+1, hi)
}

// Algorithms Fourth Edition 书上的写法
func partitionNew(a []int, lo, hi int) int {
	fmt.Printf("lo: %v, hi: %v\n", lo, hi)
	pivot := a[lo]
	i := lo
	j := hi + 1
	for true {
		fmt.Printf("i: %v, j: %v, len(a): %v\n", i, j, len(a))
		i++
		for a[i] < pivot {
			if i == hi {
				break
			}
			i++
		}
		j--
		for a[j] > pivot {
			if j == lo {
				break
			}
			j--
		}

		if i >= j {
			break
		}
		// swap
		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[j] = a[j], a[lo]
	return j
}

func partition(a []int, lo, hi int) int {
	fmt.Printf("lo: %v, hi: %v\n", lo, hi)
	pivot := a[lo]
	i := lo + 1
	j := hi
	for true {
		fmt.Printf("i: %v, j: %v, len(a): %v\n", i, j, len(a))
		for a[i] < pivot {
			// 这个判断是必须的
			// 防止数组越界
			if i == hi {
				break
			}
			i++
		}
		for a[j] > pivot {
			// 这个判断是必须的
			// 防止数组越界
			if j == lo {
				break
			}
			j--
		}

		// 这个判断必须放在这里，不能放在循环的判断条件里
		// 因为需要提前退出，否则会先执行了 swap 后再退出，那时候就已经晚了
		if i >= j {
			break
		}

		// swap
		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[j] = a[j], a[lo]
	return j
}

///////////////////////////////////////////////

func heapSort(a []int) {
	// heapify
	N := len(a)
	// 之所以从 N/2 开始，是因为后面都是大小为1的子堆，无需处理
	// 画个图就清楚了
	for k := N / 2; k >= 1; k-- {
		sink(a, k, N)
	}

	// sort
	// 先把 root 跟最后一个元素交换（因为最后一个元素是最大的元素）
	// 再对新的 root 执行 sink 操作，如此循环即可
	for N > 1 {
		swap(a, 1, N)
		N--
		sink(a, 1, N)
	}
}

// k: idx to sink
// n: number of elements in a
func sink(a []int, k, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && less(a, j, j+1) {
			j++
		}
		if less(a, j, k) {
			break
		}
		swap(a, j, k)
		k = j
	}
}

// Indices are "off-by-one" to support 1-based indexing.
// 因为这里的输入数组是从 0 开始数的
// 参考：https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/Heap.java.html
func less(a []int, i, j int) bool {
	return a[i-1] < a[j-1]
}

func swap(a []int, i, j int) {
	a[i-1], a[j-1] = a[j-1], a[i-1]
}
