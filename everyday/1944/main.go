package main

import (
	"fmt"
)

/*
有 n 个人排成一个队列，从左到右 编号为 0 到 n - 1 。给你以一个整数数组 heights ，每个整数 互不相同，heights[i] 表示第 i 个人的高度。

一个人能 看到 他右边另一个人的条件是这两人之间的所有人都比他们两人 矮 。更正式的，第 i 个人能看到第 j 个人的条件是 i < j 且 min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1]) 。

请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是第 i 个人在他右侧队列中能 看到 的 人数 。*/

// 单调栈,严格递减, 单调栈形成的过程
func canSeePersonsCount(heights []int) []int {
	ans := make([]int, len(heights))
	zhan := []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		if len(zhan) == 0 {
			zhan = append(zhan, heights[i])
			ans[i] = 0
			continue
		}
		for zhan[len(zhan)-1] < heights[i] && len(zhan) > 0 {
			zhan = zhan[:len(zhan)-1]
			ans[i]++
		}
		if len(zhan) > 0 {
			ans[i]++
		}
		zhan = append(zhan, heights[i])

	}
	return ans

}

type x struct {
	temp  int
	index int
}

// 单调栈两点性质  1.去掉的每个元素可以和当前元素形成槽 2. 栈去掉的最后元素是当前遍历元素的下一个最大元素
// 回顾一下每日温度
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	zhan := make([]x, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(zhan) == 0 { //可能比他大的为0个
			zhan = append(zhan, x{temperatures[i], i})
			ans[i] = 0
			continue
		}
		for len(zhan) > 0 && zhan[len(zhan)-1].temp <= temperatures[i] {
			zhan = zhan[:len(zhan)-1]
		}
		if len(zhan) == 0 {
			ans[i] = 0
		} else {
			ans[i] = zhan[len(zhan)-1].index - i //怎么知道坐标
		}
		zhan = append(zhan, x{temperatures[i], i})
	}
	return ans
}

func main() {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
}
