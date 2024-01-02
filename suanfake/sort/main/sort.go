package main

import "fmt"

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x IntSlice) Sort()              { Sort(x) }

// 实现了这三个接口的数据结构可以排序
func Sort(data Interface) {
	//数据长度
	n := data.Len()
	if n <= 1 {
		return
	}
	bubblesort(data, 0, n)
}
func bubblesort(data Interface, a, b int) {
	for i := a; i < b; i++ {
		for j := a; j < b-i-1; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}
func main() {
	nums := IntSlice{2, 1, 3, 4, 6, 1}
	IntSlice.Sort(nums)
	fmt.Println(nums)
}
