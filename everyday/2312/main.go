package main

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	//记忆搜索
	for _, v := range prices {
		dp[v[0]][v[1]] = v[2]
	}
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= i/2; k++ {
				dp[i][j] = max(dp[i-k][j]+dp[k][j], dp[i][j])
			}
			for k := 1; k <= j/2; k++ {
				dp[i][j] = max(dp[i][j-k]+dp[i][k], dp[i][j])
			}
		}
	}
	return int64(dp[m][n])
}
func main() {

}
