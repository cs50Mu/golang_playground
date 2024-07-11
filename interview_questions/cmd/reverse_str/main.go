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

	for i := 0; i < len(s)/2; i++ {
		chars[i], chars[len(s)-i-1] = chars[len(s)-i-1], chars[i]
	}

	return string(chars), nil
}
