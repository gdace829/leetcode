package main

const mx int = 1e5 + 1

var np = [mx]bool{1: true} //true 是非质数 false是质数
// 构建非质数数组
func init() {
	for i := 2; i*i < mx; i++ {
		if np[i] == false {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}
func countPaths(n int, edges [][]int) int64 {
	// 构建领结表
	biao := make([][]int, n+1)
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		biao[y] = append(biao[y], x)
		biao[x] = append(biao[x], y)
	}
	var nodes []int
	// 设计联通块遍历
	var dfs func(int, int)
	dfs = func(i1, fa int) {
		nodes = append(nodes, i1)
		for _, v := range biao[i1] {
			if v != fa && np[v] {
				dfs(v, i1)
			}
		}
	}
	size := make([]int, n+1)

	var res int64
	for i := 1; i <= n; i++ {
		if np[i] {
			continue
		}
		sum := 0
		for _, v := range biao[i] {
			if !np[v] {
				continue
			}

			if size[v] == 0 {
				nodes = []int{}
				dfs(v, -1)
				for _, val := range nodes {
					size[val] = len(nodes)
				}
			}
			res += int64(size[v]) * int64(sum)
			sum += size[v]
		}
		res += int64(sum)
	}
	return res

}
