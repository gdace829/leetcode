package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println((a / 2) * (a / 2))
	} else {
		fmt.Println(((a - 1) / 2) * ((a + 1) / 2))
	}
}
