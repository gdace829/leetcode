package main

const mod = 1_000_000_007

func minNonZeroProduct(p int) int {
	k := 1<<p - 1
	return (k % mod) * pow(k-1, p-1) % mod
}
func pow(x, n int) int {
	res := 1
	for ; n > 0; n-- {
		res *= (x) % mod
		x *= (x % mod)
		x %= mod
	}
	return res
}
