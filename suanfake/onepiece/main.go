package main

import "fmt"

var dp [1005][1]int

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	var n int
	var a [1005]int
	m := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if a[i] >= a[j] {
				dp[i][0] = max(dp[j][0]+a[i], dp[i][0])
				if dp[i][0] > m {
					m = dp[i][0]
				}
			}
		}
	}
	fmt.Println(m)

}
