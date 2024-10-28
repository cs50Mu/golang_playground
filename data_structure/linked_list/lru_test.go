package linked_list

import "testing"

func TestLRUCache(t *testing.T) {
	cache := ConstructorLRU(2)

	cache.Get(1)
	cache.Put(2, 2)
}
