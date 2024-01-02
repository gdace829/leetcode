package main

import (
	"fmt"
	"sort"
)

func findWisdomFruits(m int, n int, heights []int) (int, int) {
	sort.Ints(heights)
	i, j := 0, m-1
	for i < j {
		if heights[i]+heights[j] == n {
			return heights[i], heights[j]
		} else if heights[i]+heights[j] < n {
			i++
		} else {
			j--
		}
	}
	return -1, -1
}

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)
	heights := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&heights[i])
	}
	a, b := findWisdomFruits(m, n, heights)
	if a != b && a != -1 && b != -1 {
		fmt.Printf("%d %d\n", a, b)
	} else {
		fmt.Print("-1")
	}
}
