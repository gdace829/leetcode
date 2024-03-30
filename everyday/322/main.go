package main

import "math"

func coinChange(coins []int, amount int) int {
	n := len(coins)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			memo[i][j] = -1 //没有访问过
		}
	}
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	var dfs func(int, int) int
	dfs = func(i, cur int) int {
		if i < 0 {
			if cur == 0 {
				return 0
			}
			return math.MaxInt / 2
		}

		// 记忆化搜索
		p := &memo[i][cur]
		if *p != -1 {
			return *p
		}
		if coins[i] > cur {
			*p = dfs(i-1, cur)
			return *p
		}
		*p = min(dfs(i, cur-coins[i])+1, dfs(i-1, cur))
		return *p
	}

	ans := dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
