package dynamic_programming

// https://leetcode.com/problems/coin-change-ii/
// 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。

// 示例 1:

// 输入: amount = 5, coins = [1, 2, 5]
// 输出: 4
// 解释: 有四种方式可以凑成总金额:

// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1

func CoinChange2(coins []int, amt int) int {
	var dp func(idx, currAmt int) int
	dp = func(idx, currAmt int) int {
		if currAmt > amt {
			return 0
		}
		if currAmt == amt {
			return 1
		}
		if idx >= len(coins) {
			return 0
		}

		// 1. 选它，下次还能选它
		// 2. 不选它
		return dp(idx, currAmt+coins[idx]) + dp(idx+1, currAmt)
	}

	return dp(0, 0)
}
