package main

func addMinimum(word string) int {
	dp := make([]int, len(word))
	dp[0] = 2
	for i := 1; i < len(dp); i++ {
		if word[i] > word[i-1] {
			dp[i] = dp[i-1] - 1
		} else {
			dp[i] = dp[i-1] + 2
		}
	}
	return dp[len(word)-1]
}
