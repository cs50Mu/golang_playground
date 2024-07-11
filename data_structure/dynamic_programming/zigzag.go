package dynamic_programming

// https://community.topcoder.com/stat?c=problem_statement&pm=1259&rd=4493
// https://www.geeksforgeeks.org/longest-zig-zag-subsequence/
func zigZag(seq []int) int {
	// dp[i][0] 表示以 ith 元素结尾的 sequence ，而且该元素的前一个元素比它小
	// dp[i][1] 表示以 ith 元素结尾的 sequence ，而且该元素的前一个元素比它大
	// 那么递推公式可以写成：
	// dp[i][0] = max(dp[j][1] + 1)
	// dp[i][1] = max(dp[j][0] + 1)
	// 其中，0 <= j < i
	dp := make([][]int, len(seq))
	// init
	for i := 0; i < len(seq); i++ {
		dp[i] = make([]int, 2)
	}

	// base case
	dp[0][0] = 1
	dp[0][1] = 1

	for i := 1; i < len(seq); i++ {
		for j := 0; j < i; j++ {
			if seq[i] > seq[j] {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			} else if seq[i] < seq[j] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
				// ==
			} else {

			}
		}
	}

	var max int
	for i := 0; i < len(seq); i++ {
		for j := 0; j < 2; j++ {
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max
}
