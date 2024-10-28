package sorting

import (
	"container/heap"
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
	for {
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
	for {
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

// mergeSort...
// 左部分排好序，右部分排好序，最后利用merge过程让左右整体有序
// ref: https://www.bilibili.com/video/BV1wu411p7r7/
type MergeSort struct {
	arr []int
	aux []int // 辅助数组
}

func NewMergeSort(arr []int) *MergeSort {
	return &MergeSort{
		arr: arr,
		aux: make([]int, len(arr)), // 其大小等于要排序的数组的大小
	}
}

func (ms *MergeSort) Sort() {
	// ms.sortRec(0, len(ms.arr)-1)
	ms.sortIter()
}

func (ms *MergeSort) sortRec(l, r int) {
	if l == r {
		return
	}
	m := (l + r) / 2
	ms.sortRec(l, m)
	ms.sortRec(m+1, r)
	ms.merge(l, m, r)
}

func (ms *MergeSort) sortIter() {
	// step starts from 1, every loop `step`
	// becomes 2 times bigger than before
	n := len(ms.arr)
	for step := 1; step < n; step <<= 1 {
		l := 0
		for l < n {
			m := l + step - 1
			// 若右边不存在（越界了）就不用merge了
			if m+1 >= n {
				break
			}
			// m+1+x-1 = l+x-1+1+x-1 = l+2x-1
			// 计算的右边界，有可能也越界了，所有需要跟数组边界取最小值
			// 即，右半部分可能不够 step 个
			r := min(l+2*step-1, n-1)
			// merge
			// l..m..r l..m..r
			ms.merge(l, m, r)
			// 继续merge下一组 l..m..r
			l = r + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (ms *MergeSort) merge(l, m, r int) {
	i := l // index of aux array
	a := l
	b := m + 1
	for a <= m && b <= r {
		if ms.arr[a] <= ms.arr[b] {
			ms.aux[i] = ms.arr[a]
			a++
			i++
		} else {
			ms.aux[i] = ms.arr[b]
			b++
			i++
		}
	}
	for a <= m {
		ms.aux[i] = ms.arr[a]
		a++
		i++
	}
	for b <= r {
		ms.aux[i] = ms.arr[b]
		b++
		i++
	}

	// copy to arr
	for i := l; i <= r; i++ {
		ms.arr[i] = ms.aux[i]
	}
}

// https://www.bilibili.com/video/BV1cc411F7Y6
type QuickSort struct {
	arr []int
}

func NewQuickSort(arr []int) *QuickSort {
	return &QuickSort{
		arr: arr,
	}
}

func (qs *QuickSort) Sort() {
	// qs.sort(0, len(qs.arr)-1)
	qs.sortNetherlandsFlag(0, len(qs.arr)-1)
}

func (qs *QuickSort) sort(l, r int) {
	// base case
	if l >= r {
		return
	}
	// choose random x
	x := qs.arr[l+rand.Intn(r-l+1)]
	// partition by x, return the index of x
	m := qs.partition(l, r, x)
	// sort the left part and right part recursively
	// 注意：m 所在位置的元素不用再参与排序了，它的位置已经确定
	qs.sort(l, m-1)
	qs.sort(m+1, r)
}

// partition partitions arr[l..r] by x
// all the elements which are <= x are put in the left part
// all the elements which are > x are put in the right part
// finally, the x should be in the right boundary of the left part
func (qs *QuickSort) partition(l, r int, x int) int {
	var xi int // the index of the element which value is `x`
	a := l     // elements of arr[:a] are all <= x
	for i := l; i <= r; i++ {
		if qs.arr[i] <= x {
			// swap arr[i] with arr[a]
			qs.arr[i], qs.arr[a] = qs.arr[a], qs.arr[i]
			// record xi
			if qs.arr[a] == x {
				xi = a
			}
			a++
		}
	}

	// put `x` in the right position
	qs.arr[xi], qs.arr[a-1] = qs.arr[a-1], qs.arr[xi]
	return a - 1
}

// sortNetherlandsFlag 荷兰国旗问题优化
func (qs *QuickSort) sortNetherlandsFlag(l, r int) {
	if l >= r {
		return
	}

	x := qs.arr[l+rand.Intn(r-l+1)]
	a, b := qs.partitionNetherlandsFlag(l, r, x)
	qs.sortNetherlandsFlag(l, a-1)
	qs.sortNetherlandsFlag(b+1, r)
}

func (qs *QuickSort) partitionNetherlandsFlag(l, r int, x int) (int, int) {
	return partitionNetherlandFlag(qs.arr, l, r, x)
}

// partitionNetherlandsFlag 荷兰国旗问题优化
// 将 arr[l..r] 划分为三个区域：<x ==x >x
func partitionNetherlandFlag(arr []int, l, r int, x int) (int, int) {
	i := l
	a := l // the elements of [:a] are all < x
	b := r // the elements of [b+1:] are all > x
	for i <= b {
		if arr[i] < x {
			// swap
			arr[i], arr[a] = arr[a], arr[i]
			i++
			a++
		} else if arr[i] == x {
			i++
		} else {
			arr[i], arr[b] = arr[b], arr[i]
			b--
		}
	}

	// the elements of arr[a:b+1] are all ==x
	return a, b

}

// randomizedSelect 随机选择算法，时间复杂度 O(n)
// i 表示: 若 arr 排序的话，在 i 位置的数字是什么
func randomizedSelect(arr []int, i int) int {
	var ans int
	for l, r := 0, len(arr)-1; l <= r; {
		x := arr[l+rand.Intn(r-l+1)]
		a, b := partitionNetherlandFlag(arr, l, r, x)
		if a > i {
			r = a - 1
		} else if b < i {
			l = b + 1
		} else { // i is between a..b
			ans = arr[i]
			break
		}
	}
	return ans
}

// https://leetcode.com/problems/kth-largest-element-in-an-array/
// https://www.bilibili.com/video/BV1mN411b71K/
func findKthLargest(arr []int, k int) int {
	return randomizedSelect(arr, len(arr)-k)
}

// https://www.bilibili.com/video/BV1fu4y1q77y
type HeapSort struct {
	arr []int
}

func NewHeapSort(arr []int) *HeapSort {
	return &HeapSort{
		arr: arr,
	}
}

func (hs *HeapSort) Sort() {
	// hs.heapSort1()
	hs.heapSort2()
}

// 从顶到底建立大根堆， O(n*logn)
// 依次弹出堆内最大值并排好序， O(n*logn)
// 整体时间复杂度： O(n*logn)
func (hs *HeapSort) heapSort1() {
	n := len(hs.arr)
	for i := 0; i < n; i++ {
		heapInsert(hs.arr, i)
	}

	for i := n - 1; i >= 1; i-- { // 遍历到 1 即可，最后剩一个元素的时候不必排序了
		hs.arr[0], hs.arr[i] = hs.arr[i], hs.arr[0]
		heapify(hs.arr, 0, i)
	}
}

// 从底到顶建立大根堆， O(n)
// 依次弹出堆内最大值并排好序， O(n*logn)
// 整体时间复杂度： O(n*logn)
func (hs *HeapSort) heapSort2() {
	n := len(hs.arr)
	for i := n - 1; i >= 0; i-- {
		heapify(hs.arr, i, n)
	}

	for i := n - 1; i >= 1; i-- { // 遍历到 1 即可，最后剩一个元素的时候不必排序了
		hs.arr[0], hs.arr[i] = hs.arr[i], hs.arr[0]
		heapify(hs.arr, 0, i)
	}
}

// i 位置的数，向上调整 -- 大根堆
func heapInsert(arr []int, i int) {
	// 无需判断size，当到达 root 节点时，
	// arr[i] > arr[(i-1)/2] 的条件自然就不满足了
	for arr[i] > arr[(i-1)/2] {
		// swap
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
		i = (i - 1) / 2
	}
}

// i 位置的数，向下调整 -- 大根堆
func heapify(arr []int, i int, size int) {
	left := 2*i + 1 // 左节点索引
	for left < size {
		var best int                                  // 找出“更好”的节点，来判断是否需要 swap
		if left+1 < size && arr[left+1] > arr[left] { // 若有右节点且右节点“更好”
			best = left + 1
		} else {
			best = left
		}

		// 需要 swap
		if arr[best] > arr[i] {
			arr[best], arr[i] = arr[i], arr[best]
		} else {
			break // 无需 swap 时，需要立即退出
		}

		// 继续往下走
		i = best // 注意要更新 i 的值！
		left = 2*i + 1
	}
}

// Heap generic heap impl
// ref: https://go.dev/play/p/4tP6OVcKrma
// https://www.reddit.com/r/golang/comments/188it98/why_are_go_heaps_so_complicated/
type Heap[E any] struct {
	arr  []E
	size int
	// 传入 i < j 为小顶堆
	// 传入 i > j 为大顶堆
	less func(E, E) bool
}

func NewHeap[E any](less func(E, E) bool) *Heap[E] {
	return &Heap[E]{
		less: less,
	}
}

// func NewHeapWithSlice(arr []int) *Heap {
// 	return &Heap{
// 		arr:  arr,
// 		size: len(arr),
// 	}
// }

// Push add new element to `Heap`
// and rebuild the Heap
func (h *Heap[E]) Push(x E) {
	h.arr = append(h.arr, x)
	h.heapInsert(h.size)
	h.size++
}

// heapInsert i 位置的元素向上调整
func (h *Heap[E]) heapInsert(i int) {
	// 无需判断size，当到达 root 节点时，
	// arr[i] > arr[(i-1)/2] 的条件自然就不满足了
	for h.less(h.arr[i], h.arr[(i-1)/2]) {
		// swap
		h.arr[i], h.arr[(i-1)/2] = h.arr[(i-1)/2], h.arr[i]
		i = (i - 1) / 2
	}
}

// heapify i 位置的元素向下调整
func (h *Heap[E]) heapify(i int, size int) {
	left := 2*i + 1 // 左节点索引
	for left < size {
		var best int                                             // 找出“更好”的节点，来判断是否需要 swap
		if left+1 < size && h.less(h.arr[left+1], h.arr[left]) { // 若有右节点且右节点“更好”
			best = left + 1
		} else {
			best = left
		}

		// 需要 swap
		if h.less(h.arr[best], h.arr[i]) {
			h.arr[best], h.arr[i] = h.arr[i], h.arr[best]
		} else {
			break // 无需 swap 时，需要立即退出
		}

		// 继续往下走
		i = best // 注意要更新 i 的值！
		left = 2*i + 1
	}
}

// Pop remove the top element from `Heap`
// and rebuild the Heap
func (h *Heap[E]) Pop() E {
	res := h.arr[0]
	h.size--
	h.arr[0], h.arr[h.size] = h.arr[h.size], h.arr[0] // swap
	h.arr = h.arr[:h.size]                            // shrink the arr

	h.heapify(0, h.size)

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type pq []*ListNode

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	return q[i].Val < q[j].Val
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pq) Push(x any) {
	item := x.(*ListNode)
	*q = append(*q, item)
}

func (q *pq) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*q = old[0 : n-1]
	return item
}

// https://leetcode.com/problems/merge-k-sorted-lists/description/
// https://www.bilibili.com/video/BV1Gm4y1p7UE
func mergeKLists(lists []*ListNode) *ListNode {
	pq := &pq{}
	heap.Init(pq)
	for i := 0; i < len(lists); i++ {
		head := lists[i]
		if head != nil {
			heap.Push(pq, head)
		}
	}

	var res, curr *ListNode
	for pq.Len() > 0 {
		nodeI := heap.Pop(pq)
		node := nodeI.(*ListNode)
		if curr == nil {
			curr = node // save head
		} else {
			curr.Next = node
			curr = curr.Next
		}
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}

	return res
}

// https://leetcode.com/problems/minimum-operations-to-halve-array-sum/
func halveArray(nums []int) int {
	// 大顶堆
	h := NewHeap[float64](func(i, j float64) bool {
		return i > j
	})

	var sum float64
	for i := 0; i < len(nums); i++ {
		h.Push(float64(nums[i]))
		sum += float64(nums[i])
	}

	var cnt int
	target := sum / 2
	var cutCurr float64
	for cutCurr < target {
		ele := h.Pop()
		half := float64(ele) / 2
		cutCurr += half
		h.Push(half)
		cnt++
	}

	return cnt
}
