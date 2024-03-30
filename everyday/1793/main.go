package main

import "fmt"

func maximumScore(nums []int, k int) int {
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	// 左侧单调栈获取左边最邻最小值
	left := make([]int, len(nums))
	st := make([]int, 0)
	for k, v := range nums {
		for len(st) > 0 && v <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[k] = st[len(st)-1]
		} else {
			left[k] = -1
		}
		st = append(st, k)
	}
	// 右侧单调栈获取左边最邻最小值
	right := make([]int, len(nums))
	st = st[:0]
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = len(nums)
		}
		st = append(st, i)
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if left[i] < k && right[i] > k {
			fmt.Println(left[i], right[i])
			ans = max((right[i]-left[i]-1)*nums[i], ans)
		}
	}
	return ans

}
