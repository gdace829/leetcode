package main

func change(amount int, coins []int) int {
	var dfs func(int, int) int
	memo := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		memo[i] = make([]int, amount+1)
		for j := 0; j < amount+1; j++ {
			memo[i][j] = -1 //得是-1因为可能memo ij为0
		}
	}
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		x := &memo[i][c]
		if *x != -1 {
			return *x
		}
		//记忆搜索 todo
		if coins[i] > c {
			*x = dfs(i-1, c)
		} else {
			*x = dfs(i, c-coins[i]) + dfs(i-1, c)
		}
		return *x

	}
	return dfs(len(coins)-1, amount)
}
