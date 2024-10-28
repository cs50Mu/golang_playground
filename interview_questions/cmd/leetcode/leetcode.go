package main

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)

	n := len(s)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
			if (s[i] == '1' && i < len(s)-1) || (s[i] == '2' && i < len(s)-1 && s[i+1] <= '6') {
				dp[i] += dp[i+2]
			}
		}
	}

	return dp[0]
}

// dfs(s, i) 表示 the number of ways to decode of s[i...]
func dfs(s []rune, i int) int {
	// found one solution
	if i == len(s) {
		return 1
	}

	var ans int
	// 以 0 开头一定无解
	if s[i] == '0' {
		ans = 0
	} else {
		// 不是0的话，单个字符一定可以
		ans = dfs(s, i+1)
		// 两个连续字符有可能可以
		// 比如 10 23 等是可以的
		// 27 51 等就不可以, 因为只有 1-26 是合法的
		// 注意这种情况下必须保证是两个元素，即 i < lens(s) - 1
		if (s[i] == '1' && i < len(s)-1) || (s[i] == '2' && i < len(s)-1 && s[i+1] <= '6') {
			ans += dfs(s, i+2)
		}
	}

	return ans
}

func dfs2(s []rune, i int) int {
	if i == len(s) {
		return 1
	}

	var ans int
	// 单个字符的情况
	if s[i] == '0' {
		return 0
	}
	if s[i] != '*' {
		ans = dfs2(s, i+1)
	} else {
		ans = 9 * dfs2(s, i+1)
	}
	// 连续两个字符的情况
	if i+1 < len(s) {
		if s[i] != '*' && s[i+1] != '*' { // num num
			if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
				ans += dfs2(s, i+2)
			}
		} else if s[i] != '*' && s[i+1] == '*' { // num *
			switch s[i] {
			case '1':
				ans += dfs2(s, i+2) * 9
			case '2':
				ans += dfs2(s, i+2) * 6
			}
		} else if s[i] == '*' && s[i+1] != '*' { // * num
			if s[i+1] <= '6' {
				ans += dfs2(s, i+2) * 2 // '*' can be 1 or 2
			} else if s[i+1] > '6' {
				ans += dfs2(s, i+2) // 's' can only be 1
			}
		} else if s[i] == '*' && s[i+1] == '*' { // * *
			// 11 ~ 19 && 21 ~ 26
			ans += dfs2(s, i+2) * 15
		}
	}

	return ans
}

func numDecodings2(s string) int {
	dp := make([]int, len(s)+1)

	n := len(s)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		// 单个字符的情况
		if s[i] == '0' {
			dp[i] = 0
		} else {
			if s[i] != '*' {
				dp[i] = dp[i+1]
			} else {
				dp[i] = dp[i+1] * 9
			}
		}
		// 连续两个字符的情况
		if i+1 < len(s) {
			if s[i] != '*' && s[i+1] != '*' { // num num
				if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
					dp[i] += dp[i+2]
				}
			} else if s[i] != '*' && s[i+1] == '*' { // num *
				switch s[i] {
				case '1':
					dp[i] += dp[i+2] * 9
				case '2':
					dp[i] += dp[i+2] * 6
				}
			} else if s[i] == '*' && s[i+1] != '*' { // * num
				if s[i+1] <= '6' {
					dp[i] += dp[i+2] * 2
				} else if s[i+1] > '6' {
					dp[i] += dp[i+2]
				}
			} else if s[i] == '*' && s[i+1] == '*' { // * *
				// 11 ~ 19 && 21 ~ 26
				dp[i] += dp[i+2] * 15
			}
		}
		dp[i] %= 1000000007
	}

	return dp[0]
}

func uglyNumber2(n int) int {
	dp := make([]int, n+1)

	// the first ugly number is 1
	dp[1] = 1

	// i2 / i3 / i5 分别表示乘2、乘3、乘5指向的位置索引
	// 一开始指向的是第 1 个元素
	i2 := 1
	i3 := i2
	i5 := i2
	var a, b, c int
	for i := 2; i <= n; i++ {
		a = 2 * dp[i2]
		b = 3 * dp[i3]
		c = 5 * dp[i5]

		// 下一个元素是它们之中最小的
		dp[i] = min(c, min(a, b))

		// 看情况移动指针
		// 若下一个元素是用某个分支算出来的
		// 就把指针往下移动
		if a == dp[i] {
			i2++
		}
		if b == dp[i] {
			i3++
		}
		if c == dp[i] {
			i5++
		}
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestValidParen(s string) int {
	// dp[i] 表示以 s[i] 结尾的子串（s[..i]）整体 valid 的字符个数
	dp := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		// 以 '(' 结尾的子串必定不是 valid 的
		if s[i] == '(' {
			dp[i] = 0
			// 若是以 ')' 结尾，则判断 dp[i-1]
			// 即前一个位置的 dp 值
		} else {
			if i-1 >= 0 {
				// 得到以前一个字符结束的 valid 子串的长度
				steps := dp[i-1]
				// 跳过这个子串
				p := i - steps - 1
				if p >= 0 {
					// 判断这个子串的前一个字符跟 s[i] 是否能够匹配
					ch := s[p]
					if ch == '(' {
						// 能的话就更新 dp[i] 的长度
						dp[i] += steps + 2

						// 再看看前一个字符的前一个位置的 dp 值
						// 若有的话，再更新 dp[i]
						if p-1 >= 0 {
							dp[i] += dp[p-1]
						}
					}
				}
			}
		}
	}

	// 最终的答案是 dp[0..n-1] 的最大值
	max := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func findSubstringInWraproundString(s string) int {
	ss := []rune(s)
	m := make(map[rune]int)
	length := 1
	for i := 0; i < len(ss); i++ {
		if i > 0 && isAfter(ss[i], ss[i-1]) {
			length++
		} else {
			// reset length
			length = 1
		}
		// update m[ss[i]] if new val are greater
		m[ss[i]] = max(m[ss[i]], length)
	}

	var sum int
	for _, val := range m {
		sum += val
	}

	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// isAfter returns if x comes after y in ascii
// we think 'a' is after 'z'
func isAfter(x, y rune) bool {
	return (y+1-'a')%26 == x-'a'
}

func distinctSubseqII(s string) int {
	res := 1
	ss := []rune(s)
	m := make(map[rune]int)

	mod := 1000000007
	for i := 0; i < len(s); i++ {
		new := (res - m[ss[i]] + mod) % mod
		m[ss[i]] = (m[ss[i]] + new) % mod
		res = (res + new) % mod
	}

	return (res - 1 + mod) % mod
}
