package backtracking

import (
	"fmt"
)

func permutations(s string) {
	perm("", s)
}

func perm(sofar, rest string) {
	if len(rest) == 0 {
		fmt.Println(sofar)
		return
	}

	// 有循环
	// 有 n (n=len(rest)) 个递归调用
	for i := 0; i < len(rest); i++ {
		perm(sofar+rest[i:i+1], rest[:i]+rest[i+1:])
	}
}

// func subsets2(s string) {
// 	var dp func(idx int)
// 	var track []byte
// 	var res [][]byte
// 	dp = func(idx int) {
// 		res = append(res, track)
// 		if idx == len(s) {
// 			return
// 		}

// 		for i := idx; i < len(s); i++ {
// 			track = append(track, s[i])
// 			dp(i + 1)
// 			track = track[:len(track)-1]
// 		}
// 	}

// 	dp(0)

// 	fmt.Printf("subsets2: %+v\n", res)
// }

func subsets(s string) {
	ss("", s)
}

// 此算法与 permute 的区别：
// permute 的话，之前用过的还会再用
// subsets 的话，之前用过的就直接扔掉了
func ss(sofar, rest string) {
	if len(rest) == 0 {
		fmt.Println(sofar)
		return
	}

	// 这里没有循环
	// 只有两个递归调用
	// choose
	ss(sofar+rest[0:1], rest[1:])
	// not choose
	ss(sofar, rest[1:])
}
