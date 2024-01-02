package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	nums := make([]int, len(s))
	nums[0] = 1
	for i := 1; i < len(s); i++ {

		nums[i] = nums[i-1]
		num, _ := strconv.Atoi(s[i-1 : i+1])
		if num <= 36 {
			if i-2 >= 0 {
				nums[i] += nums[i-2]
			} else {
				nums[i] += 1
			}
		}

	}
	return nums[len(s)-1]
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(numDecodings(s))
}
