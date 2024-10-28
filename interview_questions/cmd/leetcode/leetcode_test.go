package main

import (
	"fmt"
	"testing"
)

// func TestDecodeWays(t *testing.T) {
// 	// s := "1*"
// 	s := "*1*1*0"

// 	target := 18
// 	// res := dfs([]rune(s), 0)
// 	res := numDecodings2(s)
// 	// res := dfs2([]rune(s), 0) % 1000000007
// 	if res != target {
// 		t.Errorf("want: %v, got: %v", target, res)
// 	}
// }

func TestUglyNumber(t *testing.T) {
	n := 1

	target := 12
	res := uglyNumber2(n)
	if res != target {
		t.Errorf("want: %v, got: %v", target, res)
	}
}

func TestLongestValidParen(t *testing.T) {
	s := "()()))))()()("

	target := 4
	res := longestValidParen(s)
	if res != target {
		t.Errorf("want: %v, got: %v", target, res)
	}
}

func TestSubStrInWrapingStr(t *testing.T) {
	s := "cdefghefghijklmnopqrstuvwxmnijklmnopqrstuvbcdefghijklmnopqrstuvwabcddefghijklfghijklmabcdefghijklmnopqrstuvwxymnopqrstuvwxyz"

	fmt.Printf("xx: %v\n", isAfter('a', 'z'))

	target := 6
	res := findSubstringInWraproundString(s)
	if res != target {
		t.Errorf("want: %v, got: %v", target, res)
	}
}

func TestDistinctSubseqII(t *testing.T) {
	s := "zchmliaqdgvwncfatcfivphddpzjkgyygueikthqzyeeiebczqbqhdytkoawkehkbizdmcnilcjjlpoeoqqoqpswtqdpvszfaksn"

	target := 3
	res := distinctSubseqII(s)
	if res != target {
		t.Errorf("want: %v, got: %v", target, res)
	}
}
