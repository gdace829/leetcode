package main

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	left, right := 1, int(2e8)
	for left <= right {
		mid := (left + right) / 2
		var valid bool
		for i := 0; i < k; i++ {
			var spend int64
			for j := 0; j < n; j++ {
				spend += max(int64(composition[i][j])*int64(mid)-int64(stock[j]), int64(0)) * int64(cost[j])
			}
			if spend <= int64(budget) {
				valid = true
				break
			}
		}
		if valid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}
