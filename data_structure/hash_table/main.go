package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	// fmt.Printf("hash: %v\n", hash("hash"))
	// fmt.Printf("it: %v\n", hash("it"))
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
		fmt.Printf("%c: %v\n", s[i], int64(s[i]))
		hash = (hash*31 + int64(s[i])) % ht.M
	}
	return hash
}

func (ht *HashTable) delete(key string) {

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
		ht.M = ht.M * 2
		ht.resize(ht.M)
	}

	var i int64
	for i = ht.hash(key); ht.Keys[i] != nil; i = (i + 1) % ht.M {
		// replace
		if key == *ht.Keys[i] {
			ht.Keys[i] = &key
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

}
