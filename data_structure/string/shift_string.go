package string

// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
// 请定义一个函数实现字符串左旋转操作的功能。
// 比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
func shiftString(s string, k int) string {
	ss := []rune(s)
	// reverse the first k (0..k-1) element
	reverse(ss, 0, k-1)
	// reverse the k..len(s)-1 element
	reverse(ss, k, len(s)-1)
	// reverse the entire string
	reverse(ss, 0, len(s)-1)

	return string(ss)
}

func reverse(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
