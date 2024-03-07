package main

// dfs
func reachableNodes(n int, edges [][]int, restricted []int) int {
	biao := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		biao[x] = append(biao[x], y)
		biao[y] = append(biao[y], x)
	}
	m := make(map[int]bool)
	for _, v := range restricted {
		m[v] = true
	}
	ans := 0
	var dfs func(int, int)
	dfs = func(i1, fa int) {
		ans++
		for _, v := range biao[i1] {
			if v != fa {
				if ok, _ := m[v]; !ok {
					dfs(v, i1)
				}

			}
		}
	}
	dfs(0, -1)
	return ans
}
