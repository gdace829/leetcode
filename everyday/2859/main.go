package main

import "fmt"

func sumIndicesWithKSetBits(nums []int, k int) int {
	res := 0
	for key, v := range nums {
		count := 0
		for key != 0 {
			if key&1 == 1 {
				count++
			}
			key >>= 1
		}
		if count == k {
			fmt.Println(v)
			res += v
		}
	}
	return res
}
