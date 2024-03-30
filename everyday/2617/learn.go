package main

import (
	"math"
	"sort"
)

func minimumVisitedCells(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	linestc := make([][]pair, n)
	for i := 0; i < n; i++ {
		linestc[i] = make([]pair, 0)
	}
	var mn int
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	rowstc := make([]pair, 0)
	//初始化行栈 列栈
	for i := m - 1; i >= 0; i-- {
		rowstc = rowstc[:0]
		for j := n - 1; j >= 0; j-- {
			d := grid[i][j]
			if i < m-1 || j < n-1 {
				mn = math.MaxInt //初始化当前值
			}
			//在行中找
			if len(rowstc) > 0 {
				k := sort.Search(len(rowstc), func(z int) bool { return rowstc[z].i <= j+d }) //第一个小于的i+d
				if k < len(rowstc) {
					mn = min(rowstc[k].x, mn)
				}
			}
			//在列中找
			if len(linestc[j]) > 0 {
				k := sort.Search(len(linestc[j]), func(z int) bool { return linestc[j][z].i <= i+d }) //第一个小于的i+d
				if k < len(linestc[j]) {
					mn = min(linestc[j][k].x, mn)
				}
			}
			//在其中找到
			if mn < math.MaxInt {
				mn++ //算上自己
				for len(rowstc) > 0 && rowstc[len(rowstc)-1].x >= mn {
					rowstc = rowstc[:len(rowstc)-1]
				}
				rowstc = append(rowstc, pair{mn, j})
				for len(linestc[j]) > 0 && linestc[j][len(linestc[j])-1].x >= mn {
					linestc[j] = linestc[j][:len(linestc[j])-1]
				}
				linestc[j] = append(linestc[j], pair{mn, i})

			}
		}
	}

	if mn >= math.MaxInt {
		return -1
	}
	return mn
}
func main() {
	minimumVisitedCells([][]int{})
}
