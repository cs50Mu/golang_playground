package main

import "fmt"

func main() {
	// fmt.Println("vim-go")
	// fmt.Printf("hash: %v\n", hash("hash"))
	// fmt.Printf("it: %v\n", hash("it"))

	ht := NewHashTable(2)
	ht.put("1", "hello")
	ht.put("2", "world")
	ht.put("3", "linuxfish")

	val1, err := ht.get("1")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	val2, err := ht.get("2")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("%v, %v\n", val1, val2)

	ht.delete("1")

	val1, err = ht.get("1")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	ht.delete("2")
	val2, err = ht.get("2")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("val2: %v\n", val2)
}

type HashTable struct {
	M    int64 // hash table size
	N    int64 // number of key-val pairs in the table
	Keys []*string
	Vals []string
}

var ErrNotFound = fmt.Errorf("not found")

func NewHashTable(m int64) *HashTable {
	return &HashTable{
		M:    m,
		Keys: make([]*string, m),
		Vals: make([]string, m),
	}
}

// hash calc hash of a string s
func (ht *HashTable) hash(s string) int64 {
	hash := int64(0)
	for i := 0; i < len(s); i++ {
		// fmt.Printf("%c: %v\n", s[i], int64(s[i]))
		hash = (hash*31 + int64(s[i])) % ht.M
	}
	return hash
}

// 不能仅仅把要删除的 key 设置为 nil ，否则会导致后面的 key 都 search 不到了
// 需要做的是：把后面的 key 全都重新 insert 一遍
// 参考：Algorithms 4th Edition by Robert Sedgewick 的 Hash Tables 一节
func (ht *HashTable) delete(key string) {
	var i int64
	for i = ht.hash(key); ht.Keys[i] != nil; i = (i + 1) % ht.M {
		if key == *ht.Keys[i] {
			ht.Keys[i] = nil
			// reinsert
			for i = (i + 1) % ht.M; ht.Keys[i] != nil; i = (i + 1) % ht.M {
				keyRedo := ht.Keys[i]
				valRedo := ht.Vals[i]
				ht.Keys[i] = nil
				ht.Vals[i] = ""
				ht.N--
				ht.put(*keyRedo, valRedo)
			}

			ht.N--
			fmt.Printf("N: %v, M: %v\n", ht.N, ht.M)
			if ht.N <= ht.M/4 {
				ht.resize(ht.M / 2)
			}
		}
	}
}

func (ht *HashTable) get(key string) (string, error) {
	var i int64
	for i = ht.hash(key); ht.Keys[i] != nil; i = (i + 1) % ht.M {
		if key == *ht.Keys[i] {
			return ht.Vals[i], nil
		}
	}
	return "", ErrNotFound
}

func (ht *HashTable) put(key, val string) {
	if ht.N >= ht.M {
		ht.resize(ht.M * 2)
	}

	var i int64
	for i = ht.hash(key); ht.Keys[i] != nil; i = (i + 1) % ht.M {
		// replace
		if key == *ht.Keys[i] {
			// ht.Keys[i] = &key
			ht.Vals[i] = val
			return
		}
	}
	// insert new
	ht.Keys[i] = &key
	ht.Vals[i] = val
	ht.N++
}

func (ht *HashTable) resize(size int64) {
	// 新建一个 HashTable
	newHT := NewHashTable(size)
	fmt.Printf("resizing to %v\n", size)

	// 把老的 HashTable 里已有的 key 全部重新插入了新的 HashTable 里
	for i := int64(0); i < ht.M; i++ {
		if ht.Keys[i] != nil {
			newHT.put(*ht.Keys[i], ht.Vals[i])
		}
	}
	ht.Keys = newHT.Keys
	ht.Vals = newHT.Vals
	ht.M = newHT.M
}
