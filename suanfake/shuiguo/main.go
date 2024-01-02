package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	nums := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&nums[i])
	}
	max := 1
	m := make(map[int]int)
	left := 0
	right := 0
	cur := 0
	for right < num {
		if m[nums[right]] == 0 {
			cur++
		}
		m[nums[right]]++
		for cur > 2 {
			m[nums[left]]--
			if m[nums[left]] == 0 {
				cur--
			}
			left++
		}
		if right-left+1 > max {
			max = right - left + 1
		}
		right++
	}
	fmt.Println(max)

}
