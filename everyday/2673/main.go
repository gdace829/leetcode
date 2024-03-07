package main

func minIncrements(n int, cost []int) int {
	var dfs func(int) int
	sum := 0
	// 接受下面子树的和,做好平衡后往上返回
	dfs = func(index int) int {
		if index > n {
			return 0
		}
		left := dfs(index * 2)
		right := dfs(index*2 + 1)
		if left > right {
			sum += left - right
			return left + cost[index-1]
		} else {
			sum += right - left
			return right + cost[index-1]
		}
	}
	dfs(1)
	return sum
}
