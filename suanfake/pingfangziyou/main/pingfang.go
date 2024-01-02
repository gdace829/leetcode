package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isSquareFree(n int) bool {
	if n%2 == 0 {
		n = n / 2
		if n%2 == 0 {
			return false
		}
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		if n%i == 0 {
			n = n / i
			if n%i == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	var T int
	fmt.Scan(&T)
	res := []int{}
	for t := 0; t < T; t++ {
		var a, b, d int
		fmt.Scan(&a, &b, &d)

		count := 0
		for i := a; i <= b; i++ {
			if isSquareFree(i) && strings.Contains(strconv.Itoa(i), strconv.Itoa(d)) {
				count++
			}
		}
		res = append(res, count)
	}
	for i := 0; i < T; i++ {
		fmt.Println(res[i])
	}
}
