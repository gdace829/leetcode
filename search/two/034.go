package two

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < 0 || left > len(nums)-1 || nums[left] != target {
		return []int{-1, -1}
	}
	first := left
	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	last := right
	return []int{first, last}
}
