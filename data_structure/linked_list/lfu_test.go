package linked_list

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)

	lfu.Put(3, 3)
	xx := lfu.Get(2)
	fmt.Printf("get(2)= %v\n", xx)

	xx = lfu.Get(3)
	fmt.Printf("get(3)= %v\n", xx)
	fmt.Printf("key2Freq: %+v\n", lfu.key2Freq)
	fmt.Printf("freq2Key: %+v\n", lfu.freq2Key)
	lfu.Put(4, 4)
	xx = lfu.Get(1)
	fmt.Printf("get(1)= %v\n", xx)
	xx = lfu.Get(3)
	fmt.Printf("get(3)= %v\n", xx)
	xx = lfu.Get(4)
	fmt.Printf("get(4)= %v\n", xx)
}
