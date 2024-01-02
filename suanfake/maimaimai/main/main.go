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

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)

	items := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&items[i][0], &items[i][1])
	}

	dp := make([]int, m+1)
	for i := 0; i < n; i++ {
		for j := m; j >= items[i][0]; j-- {
			dp[j] = max(dp[j], dp[j-items[i][0]]+items[i][1])
		}
	}

	fmt.Println(dp[m])
}
