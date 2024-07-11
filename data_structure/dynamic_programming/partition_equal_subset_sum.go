package dynamic_programming

// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// https://leetcode.com/problems/partition-equal-subset-sum/
// 思路：先求和，问题可转化为：能否选出一些元素，使得这些元素的和等于 sum(a) / 2

// 示例 1:

// 输入: [1, 5, 11, 5]
// 输出: true
// 解释: 数组可以分割成 [1, 5, 5] 和 [11].

func canPatition(a []int) bool {
	var dp func(idx, sum int) bool
	dp = func(idx, sum int) bool {
		if sum == 0 {
			return true
		}
		if sum < 0 {
			return false
		}
		if idx == len(a) {
			return false
		}

		return dp(idx+1, sum-a[idx]) || dp(idx+1, sum)
	}

	total := sum(a)
	if total%2 != 0 {
		return false
	}

	return dp(0, total/2)
}

func sum(a []int) int {
	var res int

	for _, n := range a {
		res += n
	}

	return res
}
