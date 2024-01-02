package main

import "fmt"

func pow(x int64, n int64, m int64) int64 {
	if n == 0 {
		return 1
	}
	var ans int64
	ans = 1
	for n > 0 {
		if n&1 == 1 {
			ans *= x % m
			ans %= m
			// fmt.Println(ans, (x % m), x, m)
		}
		x = (x % m) * (x % m)
		n >>= 1
	}
	return ans % m
}

func main() {
	var a, b, c int64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Println(pow(a, b, c))
}
