package main

import "math"

func countPaths(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = math.MaxInt / 2
		}
	}
	for _, v := range roads {
		x, y, z := v[0], v[1], v[2]
		g[x][y] = z
		g[y][x] = z
	}
	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt / 2
	}
	f := make([]int, n)
	f[0] = 1
	done := make([]bool, n)
	for {
		x := -1
		for k, v := range done {
			if !v && (x < 0 || dis[k] < dis[x]) {
				x = k
			}
		}
		if x == n-1 {
			return f[n-1]
		}
		done[x] = true
		for y, v := range g[x] {
			newDis := dis[x] + v
			if newDis < dis[y] {
				dis[y] = newDis
				f[y] = f[x]
			} else if newDis == dis[y] {
				f[y] = (f[x] + f[y]) % 1000000007
			}
		}

	}
}
