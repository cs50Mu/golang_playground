package dynamic_programming

import (
	"fmt"
	"math"
)

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

// 你可以认为每种硬币的数量是无限的。

// 示例 1：

// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：

// 输入：coins = [2], amount = 3
// 输出：-1

// https://betterprogramming.pub/learn-dynamic-programming-the-coin-change-problem-22a104478f50
func CoinChange(coins []int, amt int) int {
	// dp(amt int) 表示对于凑成金额 amt 需要的最少硬币个数
	var dp func(amt int) int
	memo := make(map[int]int)
	dp = func(amt int) int {
		if val, ok := memo[amt]; ok {
			return val
		}
		// base case
		if amt == 0 {
			return 0
		}
		if amt < 0 {
			// 返回 -1 用来表示此分支无解
			return -1
		}

		res := math.MaxInt
		for _, coin := range coins {
			// rem := amt - coin
			// if rem >= 0 {
			// 	res = min(res, dp(rem)+1)
			// }

			subP := dp(amt - coin)
			if subP >= 0 {
				res = min(res, subP+1)
			}

			// subP := dp(amt - coin)
			// if subP == -1 {
			// 	continue
			// }
			// res = min(res, subP+1)
		}

		// res 没有被赋新值，说明整体无解
		if res == math.MaxInt {
			res = -1
		}

		memo[amt] = res

		return memo[amt]
	}

	return dp(amt)
}

func CoinChangeWrongSolution(coins []int, amt int) int {
	var dp func(idx, currAmt, cnt int) int
	memo := make(map[string]int)
	dp = func(idx, currAmt, cnt int) int {
		key := fmt.Sprintf("%v,%v,%v", idx, currAmt, cnt)
		if val, ok := memo[key]; ok {
			return val
		}
		if currAmt > amt {
			return math.MaxInt
		}
		if currAmt == amt {
			return cnt
		}
		if idx == len(coins) {
			return math.MaxInt
		}

		memo[key] = min(dp(idx, currAmt+coins[idx], cnt+1), dp(idx+1, currAmt, cnt))
		return memo[key]
	}

	res := dp(0, 0, 0)

	if res == math.MaxInt {
		return -1
	}

	return res
}
