package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	var n, m int
	//接收nm输入
	fmt.Scan(&n)
	fmt.Scan(&m)
	nums := make([][]int, n)
	//accept input
	for i := 0; i < n; i++ {
		nums[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&nums[i][j])
		}
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if nums[i][j] == -1 {
				dp[i][j] = -1
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = nums[0][0]
				continue
			} else if i == 0 {
				dp[i][j] = nums[i][j-1] + nums[i][j]
				continue
			} else if j == 0 {
				dp[i][j] = nums[i-1][j] + nums[i][j]
				continue
			}
			if dp[i-1][j] == -1 && dp[i][j-1] == -1 {
				dp[i][j] = -1
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + nums[i][j]

		}
	}
	fmt.Println(dp[n-1][m-1])
}
