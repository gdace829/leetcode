package main

import (
	"fmt"
	"sort"
)

type Stacks [][]int

// 实现 sort.Interface 方法
func (is Stacks) Len() int {
	return len(is)
}

func (is Stacks) Less(i, j int) bool {
	return is[i][len(is[i])-1] < is[j][len(is[j])-1] // 这里可以根据需要修改比较逻辑
}

func (is Stacks) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}
func defendDroplets(n int, nums []int) int {
	nums1 := make([]int, n)
	//初始化
	for i := 0; i < n; i++ {
		nums1[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] && nums1[j]+1 > nums1[i] {
				nums1[i] = nums1[j] + 1
			}
		}
	}
	max := 0
	for i := 0; i < n; i++ {
		if nums1[i] > max {
			max = nums1[i]
		}
	}
	return max
}

func minSystems(n int, nums []int) int {
	var stacks Stacks
	stacks = make(Stacks, 0)
	for i := 0; i < n; i++ {
		if len(stacks) == 0 {
			stacks = append(stacks, []int{nums[i]})
			continue
		}
		for j := 0; j <= len(stacks); j++ {
			if j == len(stacks) {
				stacks = append(stacks, []int{nums[i]})
				sort.Sort(stacks)
				break
			}
			if nums[i] >= stacks[j][len(stacks[j])-1] {
				continue
			} else {
				stacks[j] = append(stacks[j], nums[i])
				sort.Sort(stacks)
				break
			}
		}

	}
	return len(stacks)
}

func main() {
	var time int
	fmt.Scan(&time)
	nums := make([]int, time)
	for i := 0; i < time; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(defendDroplets(len(nums), nums))
	fmt.Println(minSystems(len(nums), nums))
}
