package dynamic_programming

func minDistance(word1, word2 string) int {
	// dp[i][j] 表示 word1[0..i-1] 和 word2[0..j-1] 的最小编辑距离
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	// base case
	for i := 1; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		dp[i][0] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minT(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minT(i, j, k int) int {
	s := min(i, j)
	return min(s, k)
}
