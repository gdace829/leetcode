package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxReward(n int, matrix [][]int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = matrix[0][0]

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}

	return dp[n-1][n-1]
}

func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&nums[j])
		}
		matrix[i] = nums
	}
	fmt.Println(maxReward(n, matrix)) // 输出：21
}
