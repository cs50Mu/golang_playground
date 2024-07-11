package dynamic_programming

func lengthOfLIS(nums []int) int {
	var lis func(i int) int
	memo := make(map[int]int)

	lis = func(i int) int {
		if val, ok := memo[i]; ok {
			return val
		}
		if i == 0 {
			return 1
		}

		var max int
		var isIncr bool
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				isIncr = true
				memo[i] = lis(j)
				if memo[i] > max {
					max = memo[i]
				}
			}
		}

		if isIncr {
			memo[i] = max + 1
		} else {
			memo[i] = 1
		}
		return memo[i]
	}

	var res int
	for i := 0; i < len(nums); i++ {
		tmp := lis(i)
		if tmp > res {
			res = tmp
		}
	}

	return res
}

// dp version
func lengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))
	// dp[i]代表以ith元素结尾的字符串的lis，lis必须以ith元素结尾
	// 为什么要这么定义？可以这么考虑：因为lis必然是以某个元素结尾的
	// dp[i]初始化都为1，是因为每个元素结尾的lis至少为1（若它前面的元素都比它大时）
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		// 挨个看它前面的元素
		for j := 0; j < i; j++ {
			// 若有比它小的，则有可能
			if nums[i] > nums[j] {
				// 找出最大的
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	// 最后还是要再遍历一遍，因为最长的序列可能会以任何一个元素结尾
	var res int
	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
