package main

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		lengths := make(map[int]int)
		for _, q := range points {
			l := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			lengths[l]++
		}
		for _, v := range lengths {
			ans += v * (v - 1)
		}
	}
	return ans
}
func main() {

}
