package heap

type MaxPQ struct {
	// array index
	N int
	// underlying array
	pq []int
}

func NewMaxPQ(n int) *MaxPQ {
	return &MaxPQ{
		pq: make([]int, n+1),
	}
}

func (pq *MaxPQ) resize(m int) {
	newS := make([]int, m)

	for i := 1; i <= pq.N; i++ {
		newS[i] = pq.pq[i]
	}

	pq.pq = newS
}

func (pq *MaxPQ) Insert(v int) {
	pq.N++
	// 调整数组大小
	if pq.N == len(pq.pq)-1 {
		pq.resize(len(pq.pq) * 2)
	}
	pq.pq[pq.N] = v
	pq.swim(pq.N)
}

func (pq *MaxPQ) DelMax() int {
	res := pq.pq[1]
	pq.pq[1] = pq.pq[pq.N]
	pq.N--
	pq.sink(1)

	// 调整数组大小
	if pq.N == (len(pq.pq)-1)/4 {
		pq.resize(len(pq.pq) / 2)
	}

	return res
}

func (pq *MaxPQ) swim(k int) {
	for pq.less(k/2, k) && k > 1 {
		pq.swap(k, k/2)
		k = k / 2
	}
}

func (pq *MaxPQ) sink(k int) {
	// 为啥是 2*k?
	// 因为要朝下比较，父节点是k，则子节点就是 2*k 了，要保证子节点在范围之内
	for 2*k <= pq.N {
		j := 2 * k
		// find the bigger one of the two child
		if j < pq.N && pq.less(j, j+1) {
			j++
		}
		// 这个判断是需要的，当子节点不再大于父节点时就不需要 swap 了
		// 否则会继续向下 swap ，是不对的
		if pq.less(j, k) {
			break
		}
		pq.swap(j, k)
		k = j
	}
}

func (pq *MaxPQ) less(i, j int) bool {
	return pq.pq[i] < pq.pq[j]
}

func (pq *MaxPQ) swap(i, j int) {
	tmp := pq.pq[i]
	pq.pq[i] = pq.pq[j]
	pq.pq[j] = tmp
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
