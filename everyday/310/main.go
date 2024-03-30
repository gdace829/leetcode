package main

import "fmt"

// 树中任意节点的最长路经
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	//找到树的最长路径
	tree := make([][]int, n)
	for i := 0; i < n; i++ {
		tree[i] = make([]int, 0)
	}
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		tree[x] = append(tree[x], y)
		tree[y] = append(tree[y], x)
	}
	var path []int
	var dfs func(int, int) []int
	dfs = func(i1, fa int) []int {
		var p []int
		var res []int
		var res2 []int
		for _, v := range tree[i1] {
			if v != fa {
				p = dfs(v, i1)
				if len(p) > len(res) {
					res2 = res
					res = p
				} else if len(p) > len(res2) {
					res2 = p
				}
			}
		}
		if len(res)+len(res2)+1 > len(path) {
			fmt.Println(res, "sjs", res2)
			path = append(res, i1)
			path = append(path, res2...)
		}
		return append(res, i1)
	}
	dfs(0, -1)
	fmt.Println(path)
	d := len(path)
	if d%2 == 0 {
		return []int{path[d/2], path[d/2-1]}
	} else {
		return []int{path[d/2]}
	}
}
