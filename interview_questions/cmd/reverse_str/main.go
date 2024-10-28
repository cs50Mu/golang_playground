package main

import (
	"errors"
	"fmt"
)

// 请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，
// 翻转一个给定的字符串(可以使用单个过程变量)。

// 给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

func main() {
	input := "main"
	reversed, err := reverseStr(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("input: %v, output: %v", input, reversed)
}

var ErrMaxLenExceed = errors.New("max str len exceeded")

func reverseStr(s string) (string, error) {
	chars := []byte(s)
	if len(chars) > 5000 {
		return "", ErrMaxLenExceed
	}

	i := 0
	j := len(s) - 1
	for i < j {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}

	return string(chars), nil
}
