package dynamic_programming

// capacity: 包里最大能放多少
// profit: 每样东西值多少钱
// weight: 每样东西有多重
// n: 总共有几样东西
// https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/
func Knapsack01(capacity int, profit []int, weight []int, n int) int {
	// 没空间了或者没选择了
	if capacity == 0 || n == 0 {
		return 0
	}

	// 从最后一个开始选
	// 如果它的重量太大了放不进去，就继续选下一个
	if capacity-weight[n-1] < 0 {
		return Knapsack01(capacity, profit, weight, n-1)
		// 能放的话，就两种情况都看看（放了和没放），哪个最终的结果大就选哪个
	} else {
		return max(
			// 选它，并且下次不选它了
			Knapsack01(capacity-weight[n-1], profit, weight, n-1)+profit[n-1],
			// 不选它
			Knapsack01(capacity, profit, weight, n-1),
		)
	}
}
