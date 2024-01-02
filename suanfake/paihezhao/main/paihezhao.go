package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	count := 0
	for i, height := range heights {
		if height != expected[i] {
			count++
		}
	}

	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}
	fmt.Println(heightChecker(heights)) // 输出：3
}
