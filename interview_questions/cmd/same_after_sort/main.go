package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	s1 := "hello"
	s2 := "hlleh"

	isSame, err := isSameAfterSort(s1, s2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("res: %v\n", isSame)
}

// 给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
// 这里规定【大小写为不同字符】，且考虑字符串重点空格。
// 给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。

func isSameAfterSort(s1, s2 string) (bool, error) {
	len1, len2 := len(s1), len(s2)

	if len1 > 5000 || len2 > 5000 {
		return false, errors.New("max len exceeded")
	}

	if len1 != len2 {
		return false, nil
	}

	for _, c := range s1 {
		if strings.Count(s1, string(c)) != strings.Count(s2, string(c)) {
			return false, nil
		}
	}
	return true, nil
}
