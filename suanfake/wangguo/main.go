package main

import "fmt"

func main() {
	var n, a, b, c int64
	fmt.Scan(&n)
	// nums := make([][3]int64, n)
	for i := 0; int64(i) < n; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		min := ((b + 1) * b) / int64(2)
		max := ((a + a - b + 1) * b) / int64(2)
		if c >= min && c <= max {
			fmt.Println("Link Start")
		} else {
			fmt.Println("Attack Error")
		}
	}
	// for i := 0; i < n; i++ {
	// 	min := ((nums[i][1] + 1) * nums[i][1]) / int64(2)
	// 	max := ((nums[i][0] + nums[i][0] - nums[i][1] + 1) * nums[i][1]) / int64(2)
	// 	if nums[i][2] >= min && nums[i][2] <= max {
	// 		fmt.Println("Link Start")
	// 	} else {
	// 		fmt.Println("Attack Error")
	// 	}
	// }

}
