package main

// 每新增一个数判断以该数为最后数组是否合理
func validPartition(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = false
	if len(nums) > 1 {
		if nums[1] == nums[0] {
			dp[1] = true
		}
	}
	if len(nums) > 2 {
		if nums[2] == nums[1] && dp[1] == true {
			dp[2] = true
		}
		if nums[2]-nums[1] == 1 && nums[1]-nums[0] == 1 {
			dp[2] = true
		}
	}
	if len(nums) > 3 {
		for i := 3; i < len(nums); i++ {
			if nums[i] == nums[i-1] && dp[i-2] {
				dp[i] = true
			}
			if nums[i] == nums[i-1] && nums[i-1] == nums[i-2] && dp[i-3] {
				dp[i] = true
			}
			if nums[i]-nums[i-1] == 1 && nums[i-1]-nums[i-2] == 1 && dp[i-3] {
				dp[i] = true
			}
		}

	}
	return dp[len(dp)-1]

}
