package array

import "math"

// https://programmercarl.com/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.html#%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	var sum int
	// 先移动尾指针，当满足条件时再移动头指针，直到不满足条件
	// sum 随着头尾指针的移动在变化
	for start, end := 0, 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			// 先算距离，不然下面 start 就变了
			res = min(res, end-start+1)
			sum -= nums[start]
			start++
		}
	}

	// not found
	if res == math.MaxInt {
		return 0
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
