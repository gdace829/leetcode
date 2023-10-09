package two

func search1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[len(nums)-1] {
			if nums[mid] < target {
				if nums[mid] <= nums[len(nums)-1] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				if nums[mid] <= nums[len(nums)-1] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}

	}
	if left > len(nums)-1 || nums[left] != target {
		return -1
	}
	return left
}
