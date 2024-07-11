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
