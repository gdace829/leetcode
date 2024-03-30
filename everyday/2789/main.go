package main

func maxArrayValue(nums []int) int64 {
	sum := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if sum >= nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
	}
	return int64(sum)
}
