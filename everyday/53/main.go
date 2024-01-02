package main

import (
	"fmt"
	"math"
)

// 动态规划
func maxSubArray(nums []int) int {
	max := math.MinInt32
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max

}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
