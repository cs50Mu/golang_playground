package dynamic_programming

import (
	"fmt"
	"sort"
)

func maxEnvlopes(evps [][]int) int {
	// sort
	// first sort by w, asc, then sort by h, desc
	sort.Slice(evps, func(i, j int) bool {
		if evps[i][0] == evps[j][0] {
			return evps[i][1] > evps[j][1]
		} else {
			return evps[i][0] < evps[j][0]
		}
	})

	// solve lis using h, that's the answer
	fmt.Printf("%+v\n", evps)
	var ws []int
	for i := 0; i < len(evps); i++ {
		ws = append(ws, evps[i][1])
	}

	return lengthOfLISDP(ws)
}
