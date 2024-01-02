package main

import (
	"fmt"
)

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func dfs(poster []string, x int, y int, index int) int {
	if index == 4 {
		return 1
	}
	count := 0
	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && nx < len(poster) && ny >= 0 && ny < len(poster[0]) && poster[nx][ny:ny+1] == "CSGO"[index:index+1] {
			count += dfs(poster, nx, ny, index+1)
		}
	}
	return count
}

func main() {
	// 输入
	var n int
	fmt.Scan(&n)
	poster := make([]string, n)
	for i := range poster {
		fmt.Scan(&poster[i])
	}
	// 处理
	count := 0
	for i := range poster {
		for j := range poster[i] {
			if poster[i][j:j+1] == "C" {
				count += dfs(poster, i, j, 1)
			}
		}
	}

	// 输出
	fmt.Println(count)
}
