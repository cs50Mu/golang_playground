package delay_queue

// 学习 generic 的比较好的资料:
// https://blog.logrocket.com/understanding-generics-go-1-18/

const (
	PQTypeMax = 1
	PQTypeMin = 2
)

type Delayed interface {
	// this < other return true
	// this > other return false
	CompareTo(other Delayed) bool
}

// 为啥要用 Ordered
// https://stackoverflow.com/questions/70562572/in-go-generics-why-cant-i-use-comparable-constraint-with-order-operators/70562597#70562597
type PQ[T Delayed] struct {
	// array index
	N int
	// underlying array
	pq          []T
	compareFunc func(i, j T) bool
}

func NewPQ[T Delayed](n int, pqType int) *PQ[T] {
	var compareFunc func(i, j T) bool
	switch pqType {
	case PQTypeMax:
		compareFunc = func(i, j T) bool { return i.CompareTo(j) }
	case PQTypeMin:
		compareFunc = func(i, j T) bool { return j.CompareTo(i) }
	}
	return &PQ[T]{
		pq:          make([]T, n+1),
		compareFunc: compareFunc,
	}
}

func (pq *PQ[T]) resize(m int) {
	newS := make([]T, m)

	for i := 1; i <= pq.N; i++ {
		newS[i] = pq.pq[i]
	}

	pq.pq = newS
}

func (pq *PQ[T]) Insert(v T) {
	pq.N++
	// 调整数组大小
	if pq.N == len(pq.pq)-1 {
		pq.resize(len(pq.pq) * 2)
	}
	pq.pq[pq.N] = v
	pq.swim(pq.N)
}

func (pq *PQ[T]) DelTop() T {
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

func (pq *PQ[T]) swim(k int) {
	for pq.compare(k/2, k) && k > 1 {
		pq.swap(k, k/2)
		k = k / 2
	}
}

func (pq *PQ[T]) sink(k int) {
	// 为啥是 2*k?
	// 因为要朝下比较，父节点是k，则子节点就是 2*k 了，要保证子节点在范围之内
	for 2*k <= pq.N {
		j := 2 * k
		// find the bigger one of the two child
		if j < pq.N && pq.compare(j, j+1) {
			j++
		}
		// 这个判断是需要的，当子节点不再大于父节点时就不需要 swap 了
		// 否则会继续向下 swap ，是不对的
		if pq.compare(j, k) {
			break
		}
		pq.swap(j, k)
		k = j
	}
}

func (pq *PQ[T]) compare(i, j int) bool {
	return pq.compareFunc(pq.pq[i], pq.pq[j])
}

// func (pq *PQ) less(i, j int) bool {
// 	return pq.pq[i] < pq.pq[j]
// }

// func (pq *PQ) greater(i, j int) bool {
// 	return pq.pq[i] > pq.pq[j]
// }

func (pq *PQ[T]) swap(i, j int) {
	tmp := pq.pq[i]
	pq.pq[i] = pq.pq[j]
	pq.pq[j] = tmp
}

type DelayQueue struct {
}
