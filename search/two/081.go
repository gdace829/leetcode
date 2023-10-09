package two

func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	if target == nums[right] {
		return true
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if target > nums[len(nums)-1] {
			if nums[mid] < target {
				if nums[mid] < nums[len(nums)-1] {
					right = mid - 1
				} else if nums[mid] > nums[len(nums)-1] {
					left = mid + 1
				} else {
					z := mid
					for ; z <= len(nums)-1; z++ {
						if nums[z] != nums[len(nums)-1] {
							left = mid + 1
							break
						}
					}
					if z == len(nums) {
						right = mid - 1
					}
				}
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				if nums[mid] < nums[len(nums)-1] {
					right = mid - 1
				} else if nums[mid] > nums[len(nums)-1] {
					left = mid + 1
				} else {
					z := mid
					for ; z <= len(nums)-1; z++ {
						if nums[z] != nums[len(nums)-1] {
							left = mid + 1
							break
						}
					}
					if z == len(nums) {
						right = mid - 1
					}
				}
			}
		}

	}
	if left > len(nums)-1 || nums[left] != target {
		return false
	}
	return true
}
