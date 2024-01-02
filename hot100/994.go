package main

import "fmt"

// 定义一个结构体来表示橘子的位置
type Position struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	// dx 和 dy 分别表示上下左右四个方向
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	queue := []Position{} // 存储腐烂橘子的队列
	fresh := 0            // 新鲜橘子的数量

	// 遍历网格，统计新鲜橘子的数量并将腐烂橘子的位置加入队列
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, Position{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	// 如果没有新鲜橘子，直接返回 0
	if fresh == 0 {
		return 0
	}

	minutes := -1 // 经过的分钟数

	for len(queue) > 0 {
		minutes++
		qlen := len(queue)

		for i := 0; i < qlen; i++ {
			pos := queue[0]
			queue = queue[1:]

			for j := 0; j < 4; j++ {
				x, y := pos.x+dx[j], pos.y+dy[j]

				if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					queue = append(queue, Position{x, y})
				}
			}
		}
	}

	// 如果还有新鲜橘子没有被腐烂，返回 -1
	if fresh > 0 {
		return -1
	}

	return minutes
}

func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
