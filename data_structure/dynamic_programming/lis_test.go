package dynamic_programming

import (
	"fmt"
	"testing"
)

func TestLIS(t *testing.T) {
	nums := []int{4, 10, 4, 3, 8, 9}

	fmt.Println(lengthOfLIS(nums))
}
