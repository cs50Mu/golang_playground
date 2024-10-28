package linked_list

import "container/list"

type kv struct {
	key int
	val int
}

type LRUCache struct {
	Cap  int                   // the capicity of the cache
	size int                   // the current size of the cache
	dll  *list.List            // the underlining doubly-linked list
	m    map[int]*list.Element // the underlining hash map
}

func ConstructorLRU(cap int) LRUCache {
	return LRUCache{
		Cap: cap,
		dll: list.New(),
		m:   make(map[int]*list.Element),
	}
}

func (lru *LRUCache) Get(key int) int {
	if ele, ok := lru.m[key]; ok {
		lru.dll.MoveToFront(ele)
		kv := ele.Value.(*kv)
		return kv.val
	}
	return -1
}

func (lru *LRUCache) Put(key, val int) {
	if ele, ok := lru.m[key]; ok {
		// upadte val, no need to evict
		ele.Value.(*kv).val = val
		lru.dll.MoveToFront(ele)
	} else {
		// this is new, need to add
		// check if there is space
		if lru.size == lru.Cap {
			lru.evict()
		}
		xx := &kv{key, val}
		e := lru.dll.PushFront(xx)
		lru.m[key] = e
		lru.size++
	}
}

func (lru *LRUCache) evict() {
	dll := lru.dll
	back := dll.Back()
	kv := back.Value.(*kv)
	dll.Remove(back)
	delete(lru.m, kv.key)
	lru.size--
}
