func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] ^= 1
			}
		}
	}
	res := 0
	for j := 0; j < n; j++ {
		cnt := 0
		for i := 0; i < m; i++ {
			cnt += grid[i][j]
		}
		res += max(cnt, m-cnt) * (1 << (n - j - 1))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}