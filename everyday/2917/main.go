package main

func findKOr(nums []int, k int) int {
	cur := 1
	n := len(nums)
	res := 0
	for i := 1; i <= 32; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if cur&nums[j] == cur {
				count++
			}
		}
		if count >= k {
			res += cur
		}
		cur *= 2
	}
	return res
}
