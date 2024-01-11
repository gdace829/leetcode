package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(maximumRows([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}, 2))
}
func maximumRows(matrix [][]int, numSelect int) int {
	hang, lie := len(matrix), len(matrix[0])
	nums := make([]int, hang) //每一行的二进制转换的数
	for i := 0; i < hang; i++ {
		for j := 0; j < lie; j++ {
			nums[i] += matrix[i][j] << (lie - 1 - j)
		}
	}
	max := 0
	limit := 1 << lie
	for i := 1; i < limit; i++ {
		if bits.OnesCount(uint(i)) != numSelect {
			continue
		}
		t := 0
		for j := 0; j < hang; j++ {
			if (i & nums[j]) == nums[j] {
				t++
			}
		}
		if t > max {
			max = t
		}
	}
	return max

}
