package main

func divisibilityArray(word string, m int) []int {
	x := 0
	ans := make([]int, len(word))
	for k, v := range word {
		x = (10*x + int(v-'0')) % m
		if x == 0 {
			ans[k] = 1
		}
	}
	return ans
}
