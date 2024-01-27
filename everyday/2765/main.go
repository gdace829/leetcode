package main

func alternatingSubarray(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	max := 0
	for i := 1; i < len(nums); i++ {

		if nums[i]-nums[i-1] == -1 {
			if i >= 2 && nums[i-1]-nums[i-2] == 1 {
				dp[i] = dp[i-1] + 1
			}
		} else if nums[i]-nums[i-1] == 1 {
			if i-1 == 0 || nums[i-1]-nums[i-2] == -1 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = 2
			}
		}
		if dp[i] > max {
			max = dp[i]
		}

	}
	if max == 1 {
		return -1
	} else {
		return max
	}
}
