package main

import "fmt"

func gcd(a, b int) int {
	for a%b != 0 {
		t := a
		a = b
		b = t % b
	}
	return b
}
func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity > jug1Capacity+jug2Capacity {
		return false
	}
	x := gcd(jug1Capacity, jug2Capacity)
	if targetCapacity%x == 0 {
		return true
	}
	return false
}
func main() {
	fmt.Println(gcd(5, 10))
}
