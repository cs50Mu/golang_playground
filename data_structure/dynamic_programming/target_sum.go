package dynamic_programming

// https://leetcode.com/problems/target-sum/
// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

// 示例：

// 输入：nums: [1, 1, 1, 1, 1], S: 3
// 输出：5

func TargetSum(nums []int, target int) int {
	var dp func(idx, sum int)
	var res int
	dp = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				res++
			}
			return
		}

		dp(idx+1, sum+nums[idx])
		dp(idx+1, sum-nums[idx])
	}

	dp(0, 0)

	return res
}
