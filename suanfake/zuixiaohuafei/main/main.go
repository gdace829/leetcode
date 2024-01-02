package main

import (
	"fmt"
)

func minCost(n int, A []int) int {
	cost := 0
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			A[n-1] += A[i]
			cost -= A[i]
			A[i] = 0
		} else if A[i] > 0 {
			A[i+1] += A[i]
			A[i] = 0
		}

	}
	return cost

}

func main() {
	var n int
	var A []int
	fmt.Scan(&n)
	A = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	fmt.Println(minCost(n, A)) // 输出：3
}
