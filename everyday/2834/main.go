package main

func minimumPossibleSum(n int, target int) int {
	min := 0
	if target/2 > n {
		min = n
	} else {
		min = target / 2
	}
	return (min*(min+1)/2%1000000007 + (target+target+n-min)*(n-min)/2%1000000007) % 1000000007
}
