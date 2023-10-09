package two

func searchMatrix(matrix [][]int, target int) bool {
	x := len(matrix) - 1
	y := 0
	left := 0
	for left <= x {
		mid := left + (x-left)/2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] < target {
			left = mid + 1
		} else {
			x = mid - 1
		}
	}
	if x < 0 {
		return false
	}
	if matrix[x][y] == target {
		return true
	} else if matrix[x][y] < target {
		left := 0
		right := len(matrix[0]) - 1
		for left <= right {
			mid := left + (right-left)/2
			if matrix[x][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left > len(matrix[0])-1 {
			return false
		}
		return matrix[x][left] == target
	}

	return false

}
