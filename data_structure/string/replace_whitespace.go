package string

// s = "We are happy."
func replaceWhitespace(s string) string {
	// ss := []rune(s)

	var res []rune
	for _, c := range s {
		if c == ' ' {
			res = append(res, []rune{'%', '2', '0'}...)
		} else {
			res = append(res, c)
		}
	}

	return string(res)
}
