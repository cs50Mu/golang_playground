package dynamic_programming

import (
	"math"
)

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// 示例 1：
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

func BestTimeBuySellBruteForce(prices []int) int {
	res := math.MinInt
	for i := 0; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			diff := prices[i] - prices[j]
			if diff > 0 {
				res = max(res, diff)
			}
		}
	}

	return res
}

func BestTimeBuySell(prices []int) int {
	minPrice := prices[0]
	maxDiff := math.MinInt
	for i := 1; i < len(prices); i++ {
		// 维护一个 minPrice 变量和一个 maxDiff 变量
		// 则在遇到一个新的元素后，只可能有两种情况：
		// 1. 这个元素比 minPrice 还小，那么就更新 minPrice 变量
		// 2. 这个元素比 minPrice 大，那么就更新 maxDiff 变量
		// 因为是从左向右遍历的，所以能够保证 prices[i] 一定在 minPrice 后面出现
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxDiff = max(maxDiff, prices[i]-minPrice)
		}
	}

	return maxDiff
}

// 双指针解法
// sell 比 buy 小的时候，同时更新两个指针
// 当 sell 比 buy 大的时候，只更新 sell 指针，同时也更新 maxDiff
func BestTimeBuySellTwoPointer(prices []int) int {
	buy := 0
	sell := 1

	maxDiff := 0
	for buy < len(prices) && sell < len(prices) {
		if prices[buy] > prices[sell] {
			buy = sell
			sell = sell + 1
		} else {
			maxDiff = max(maxDiff, prices[sell]-prices[buy])
			// only move sell pointer
			sell += 1
		}
	}

	return maxDiff
}
