package main

import "strconv"

func maximumSwap(num int) int {
	nums := []byte(strconv.Itoa(num))
	n := len(nums)
	maxIndex := n - 1
	in1, in2 := -1, -1
	for i := n - 1; i >= 0; i-- {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		} else if nums[i] < nums[maxIndex] {
			in1 = i
			in2 = maxIndex

		}
	}
	if in1 < 0 {
		return num
	}
	nums[in1], nums[in2] = nums[in2], nums[in1]
	v, _ := strconv.Atoi(string(nums))
	return v
}
