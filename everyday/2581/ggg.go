package main

func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges)
	//建图
	graph := make([][]int, n+1)
	for _, v := range edges {
		x, y := v[0], v[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	//转换哈希表
	type pair struct{ x, y int }
	m := make(map[pair]int, len(guesses))
	for _, v := range guesses {
		m[pair{v[0], v[1]}] = 1
	}
	cnt := make([]int, n+1)
	var dfs func(int, int)
	dfs = func(i1, fa int) {
		for _, v := range graph[i1] {
			if v != fa {
				if _, ok := m[pair{i1, v}]; ok {
					cnt[0]++
				}
				dfs(v, i1)
			}
		}
	}
	ans := 0
	//算出cnt[0]值
	dfs(0, -1)
	//换根dp
	var reroot func(int, int, int)
	reroot = func(i1, fa, cntt int) {
		if cntt >= k {
			ans++
		}
		for _, v := range graph[i1] {
			if v != fa {
				cnt[v] = cntt + m[pair{v, i1}] - m[pair{i1, v}]

				reroot(v, i1, cnt[v])
			}
		}
	}
	reroot(0, -1, cnt[0])
	return ans
}
