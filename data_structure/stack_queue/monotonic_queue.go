package stack_queue

// MonotonicQueue 单调队列
// 这里实现的是一个单调递减队列
// 即，队首是最大值
type MonotonicQueue struct {
	// 按理说，底层应该是一个 dequeue
	// 但 Golang 标准库很简陋，万物用 slice 来模拟
	a []int
}

func NewMQ() *MonotonicQueue {
	return &MonotonicQueue{
		a: make([]int, 0),
	}
}

// Push 只要有比当前要 push 的元素小的元素，就把它们
// 从队列里删掉，最后再把当前元素 push 进去
func (mq *MonotonicQueue) Push(x int) {
	for len(mq.a) > 0 && x > mq.a[len(mq.a)-1] {
		mq.a = mq.a[:len(mq.a)-1]
	}
	mq.a = append(mq.a, x)
}

// Pop 只有当当前元素等于队首元素时才把它从队首删掉
func (mq *MonotonicQueue) Pop(x int) {
	if len(mq.a) > 0 && mq.a[0] == x {
		mq.a = mq.a[1:]
	}
}

// Front 返回的是当前队列的最大值
func (mq *MonotonicQueue) Front() int {
	return mq.a[0]
}
