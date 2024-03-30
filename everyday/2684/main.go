package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxMoves(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	dpb := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dpb[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 0
		dpb[i][0] = true
	}

	for j := 1; j < n; j++ {
		for i := 0; i < m; i++ {
			if i-1 >= 0 && j-1 >= 0 {
				if grid[i-1][j-1] < grid[i][j] && dpb[i-1][j-1] {
					dp[i][j] = max(dp[i-1][j-1]+1, dp[i][j])
					dpb[i][j] = true
				}
			}
			if i >= 0 && j-1 >= 0 {
				if grid[i][j-1] < grid[i][j] && dpb[i][j-1] {
					dp[i][j] = max(dp[i][j-1]+1, dp[i][j])
					dpb[i][j] = true
				}
			}
			if i+1 < m && j-1 >= 0 {
				if grid[i+1][j-1] < grid[i][j] && dpb[i+1][j-1] {
					dp[i][j] = max(dp[i+1][j-1]+1, dp[i][j])
					dpb[i][j] = true
				}
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
