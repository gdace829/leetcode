package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("sjs")
}
func splitArray(nums []int, k int) int {
	n := len(nums)
	f := make([][]int, n+1)
	sub := make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, k+1)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	f[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < min(i, k); j++ {
			for k := 0; k < i; k++ {
				f[i][j] = min(f[i][j], max(f[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return f[n][k]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
