package array

func generateMatrix(n int) [][]int {
	// init
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	num := 1

	// 先把一圈画出来，左闭右开
	// 再找规律，抽象出 startX 和 startY
	// 注意 startX 和 startY 一个增加一个减少
	for startX, startY := 0, n-1; startX < startY; startX, startY = startX+1, startY-1 {
		for i := startX; i < startY; i++ {
			res[startX][i] = num
			num++
		}
		for i := startX; i < startY; i++ {
			res[i][startY] = num
			num++
		}
		for i := startX; i < startY; i++ {
			res[startY][n-i-1] = num
			num++
		}
		for i := startX; i < startY; i++ {
			res[n-i-1][startX] = num
			num++
		}
	}

	// 若总数（n*n）是个奇数，那最后还得把中间的那个填上
	if (n*n)%2 == 1 {
		mid := n / 2
		res[mid][mid] = num
	}

	return res
}
