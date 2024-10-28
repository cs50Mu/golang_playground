package systemdesign

import (
	"fmt"
	"math"
	"math/rand"
)

// https://www.bilibili.com/video/BV1nF411y7rD/

// https://leetcode.com/problems/insert-delete-getrandom-o1/
type RandomizedSet struct {
	m    map[int]int // val --> index of the val in arr
	arr  []int
	size int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:   make(map[int]int),
		arr: make([]int, 200000),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.arr[this.size] = val
		this.m[val] = this.size
		this.size++
		fmt.Printf("after insert: %v, eles: %+v\n", val, this.arr[:this.size])
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.m[val]; ok {
		fmt.Printf("val: %v, idx: %v\n", val, idx)
		if idx != this.size-1 {
			endVal := this.arr[this.size-1] // the value of the last element
			// swap it with the last element
			this.arr[idx], this.arr[this.size-1] = this.arr[this.size-1], this.arr[idx]
			// don't forget to update the index of the last element in map
			this.m[endVal] = idx
		}
		delete(this.m, val)
		this.size--
		fmt.Printf("after remove: %v, eles: %+v\n", val, this.arr[:this.size])
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	fmt.Printf("eles: %+v\n", this.arr[:this.size])
	ranIdx := rand.Intn(this.size)
	return this.arr[ranIdx]
}

// https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/
type RandomizedCollection struct {
	// m 里放的是 val 对应的 idx 集合
	m map[int]map[int]bool // {val1: (1,4), val2: (2,6)}
	// val 在 arr 里是可以重复的
	arr  []int
	size int // the size of the arr
}

func Constructor2() RandomizedCollection {
	return RandomizedCollection{
		m:   make(map[int]map[int]bool),
		arr: make([]int, 200000),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	this.arr[this.size] = val
	if set, ok := this.m[val]; ok {
		set[this.size] = true
		this.size++
		return false
	}
	s := map[int]bool{this.size: true}
	this.m[val] = s
	this.size++

	return true
}

func (this *RandomizedCollection) Remove(val int) bool {
	if set, ok := this.m[val]; ok {
		var idx int
		for k := range set {
			idx = k
			break
		}
		delete(set, idx)
		if len(set) == 0 {
			delete(this.m, val)
		}

		if idx != this.size-1 {
			endVal := this.arr[this.size-1]
			this.arr[idx], this.arr[this.size-1] = this.arr[this.size-1], this.arr[idx]
			endValSet := this.m[endVal]
			// remove old index, add new index
			delete(endValSet, this.size-1)
			endValSet[idx] = true
		}

		this.size -= 1
		return true
	} else {
		return false
	}
}

func (this *RandomizedCollection) GetRandom() int {
	ranIdx := rand.Intn(this.size)
	return this.arr[ranIdx]
}

// https://leetcode.com/problems/maximum-frequency-stack/
// https://www.bilibili.com/video/BV1nF411y7rD/
type FreqStack struct {
	// m 保存的是频次出现的历史,同一个数字可能出现在不同的频次历史中
	m       map[int]*[]int // 3:[1,2] 表示出现3次的数字有 1 和 2
	freqM   map[int]int    // 频次映射表，3:5，表示数字3出现了5次
	maxFreq int            // 当前最大频率
}

func Constructor4() FreqStack {
	return FreqStack{
		m:     make(map[int]*[]int),
		freqM: make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	// 先更新频次表
	var newFreq int
	if freq, ok := this.freqM[val]; ok {
		newFreq = freq + 1
	} else {
		newFreq = 1
	}
	this.freqM[val] = newFreq

	// 再更新最大频次
	if newFreq > this.maxFreq {
		this.maxFreq = newFreq
	}

	// 最后更新频次对应的元素
	if elems, ok := this.m[newFreq]; ok {
		*elems = append(*elems, val)
	} else {
		this.m[newFreq] = &([]int{val})
	}
	// fmt.Printf("Push, maxFreq: %v\n", this.maxFreq)
}

func (this *FreqStack) Pop() int {
	// 根据maxFreq找到当前应该返回的元素，并移除
	maxFreq := this.maxFreq
	// fmt.Printf("Pop, maxFreq: %v\n", maxFreq)
	elems := this.m[maxFreq]
	elemsLen := len(*elems)
	e := (*elems)[elemsLen-1]
	*elems = (*elems)[:elemsLen-1] // remove last element

	// 更新 maxFreq
	if len(*elems) == 0 {
		delete(this.m, maxFreq)
		// also need to decrease maxFreq
		this.maxFreq -= 1
	}

	// 更新频次表
	this.freqM[e] -= 1 // decrease freq of this element
	if this.freqM[e] == 0 {
		delete(this.freqM, e)
	}

	return e
}

// https://leetcode.com/problems/all-oone-data-structure/
type AllOne struct {
	head *bucket
	tail *bucket
	m    map[string]*bucket // 某个 key 对应的 bucket
}

func Constructor6() AllOne {
	head := newBucket("", 0)
	tail := newBucket("", math.MaxInt32)
	// 记住要把 head 和 tail 先连起来才行
	head.next = tail
	tail.prev = head
	return AllOne{
		// dummy nodes
		head: head,
		tail: tail,
		m:    make(map[string]*bucket),
	}
}

func (this *AllOne) Inc(key string) {
	if bucket, ok := this.m[key]; ok { // key已经存在
		next := bucket.next
		if next.freq == bucket.freq+1 { // freq+1 的bucket刚好存在
			next.keys[key] = true
			this.m[key] = next
		} else { // 不存在就要新建并插入
			b := newBucket(key, bucket.freq+1)
			insertAfter(bucket, b)
			this.m[key] = b
		}
		// 删掉老 bucket 中的 key
		delete(bucket.keys, key)
		if len(bucket.keys) == 0 {
			remove(bucket)
		}
	} else {
		// 有可能 freq=1 的 bucket 已经有了
		if this.head.next.freq == 1 {
			this.head.next.keys[key] = true
			this.m[key] = this.head.next
		} else {
			b := newBucket(key, 1)
			insertAfter(this.head, b)
			this.m[key] = b
		}
	}
}

func (this *AllOne) Dec(key string) {
	if bucket, ok := this.m[key]; ok {
		// 当 key 的 freq 已经是 1 的时候，应该在 this.m 中删除 key
		if bucket.freq == 1 {
			delete(this.m, key)
		} else {
			prev := bucket.prev
			if prev.freq == bucket.freq-1 {
				this.m[key] = prev // update this.m[key]
				prev.keys[key] = true
			} else {
				b := newBucket(key, bucket.freq-1)
				insertAfter(bucket.prev, b)
				this.m[key] = b
			}
		}
		delete(bucket.keys, key)
		if len(bucket.keys) == 0 {
			remove(bucket)
		}
	}
	// 根据题意，bucket 找不到的情况不存在
}

func (this *AllOne) GetMaxKey() string {
	var res string
	for k := range this.tail.prev.keys {
		res = k
		break
	}
	return res
}

func (this *AllOne) GetMinKey() string {
	var res string
	for k := range this.head.next.keys {
		res = k
		break
	}
	return res
}

type bucket struct {
	keys map[string]bool
	freq int
	prev *bucket
	next *bucket
}

func newBucket(key string, freq int) *bucket {
	return &bucket{
		keys: map[string]bool{key: true},
		freq: freq,
	}
}

// insertAfter insert `toBeIncerted` after `bucket`
func insertAfter(bucket *bucket, toBeIncerted *bucket) {
	next := bucket.next
	bucket.next = toBeIncerted
	toBeIncerted.prev = bucket
	toBeIncerted.next = next
	next.prev = toBeIncerted
}

// remove removes `bucket` from the doubly-linked list
func remove(bucket *bucket) {
	bucket.prev.next = bucket.next
	bucket.next.prev = bucket.prev
}
