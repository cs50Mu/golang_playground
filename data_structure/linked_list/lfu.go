package linked_list

import (
	"fmt"
	"math"
)

type LFUCache struct {
	key2Val  map[int]int
	minFreq  int
	key2Freq map[int]int
	freq2Key map[int]*linkedHashSet
	size     int
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		key2Val:  make(map[int]int),
		key2Freq: make(map[int]int),
		freq2Key: make(map[int]*linkedHashSet),
		capacity: capacity,
		minFreq:  math.MaxInt,
	}
}

func (this *LFUCache) Get(key int) int {
	if val, ok := this.key2Val[key]; ok {
		// add freq
		this.addFreq(key)
		return val
	}
	return -1
}

func (this *LFUCache) addFreq(key int) {
	// add freq of key
	origFreq := this.key2Freq[key]
	this.key2Freq[key] += 1
	keys := this.freq2Key[origFreq]
	// delete key from orig freq
	keys.Delete(key)
	if keys.Empty() {
		// 理解这里是关键：等于这个 freq 的 key 删完了
		// 而且呢这个 freq 正好是当前最小的 freq
		// 那么 minFreq += 1 （因为刚从旧的 freq 加了1得到了新的 freq）
		// 最新的 minFreq 一定是这个值了
		if origFreq == this.minFreq {
			this.minFreq += 1
		}
	}
	// add key to new freq
	newFreq := origFreq + 1
	if newFreq < this.minFreq {
		this.minFreq = newFreq
	}
	if lhs, ok := this.freq2Key[newFreq]; ok {
		lhs.PushBack(key)
	} else {
		this.freq2Key[newFreq] = NewLHS()
		this.freq2Key[newFreq].PushBack(key)
	}
}

func (this *LFUCache) removeLFU() {
	fmt.Printf("minFreq: %v\n", this.minFreq)
	keys := this.freq2Key[this.minFreq]
	node := keys.PopFront()
	delete(this.key2Freq, node.Val)
	delete(this.key2Val, node.Val)
	this.minFreq += 1
}

func (this *LFUCache) Put(key int, value int) {
	if _, ok := this.key2Val[key]; ok {
		this.key2Val[key] = value
		// add freq of key
		this.addFreq(key)
	} else {
		// check size of cache
		// if cache is full, delete the least frequently used kv
		if this.size == this.capacity {
			this.removeLFU()
			this.size--
		}
		// then add the new kv
		this.key2Freq[key] = 1
		this.key2Val[key] = value
		// TODO: duplicated logic
		if lhs, ok := this.freq2Key[1]; ok {
			lhs.PushBack(key)
		} else {
			this.freq2Key[1] = NewLHS()
			this.freq2Key[1].PushBack(key)
		}
		if this.minFreq > 1 {
			this.minFreq = 1
		}
		this.size++
	}
}

type linkedHashSet struct {
	m    map[int]*Node
	dll  *DoublyLinkedList
	size int
}

func NewLHS() *linkedHashSet {
	return &linkedHashSet{
		m:   make(map[int]*Node),
		dll: NewDLL(),
	}
}

func (lhs *linkedHashSet) PushBack(key int) {
	node := &Node{Val: key}
	lhs.m[key] = node
	lhs.dll.PushBack(node)
	lhs.size++
}

func (lhs *linkedHashSet) PopFront() *Node {
	front := lhs.dll.PopFront()
	delete(lhs.m, front.Val)
	lhs.size--
	return front
}

func (lhs *linkedHashSet) Delete(key int) {
	node := lhs.m[key]
	lhs.dll.RemoveAt(node)
	delete(lhs.m, key)
	lhs.size--
}

func (lhs *linkedHashSet) Empty() bool {
	return lhs.size == 0
}
