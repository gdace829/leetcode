package main

import (
	"fmt"
	"sort"
)

// 间隔
type Interval struct {
	start, end int
}

func coverMuddyRoad(n, L int, intervals []Interval) int {
	//对间隔进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	count := 0
	end := 0

	for i := 0; i < n; i++ {
		//已经铺完
		if intervals[i].end <= end {
			continue
		}

		if intervals[i].start > end { //和已经铺的没有无重叠情况
			count1 := (intervals[i].end - intervals[i].start + L - 1) / L //有余的在后面补一块
			end = intervals[i].start + count1*L
			count += count1
		} else { //有重叠
			count1 := (intervals[i].end - end + L - 1) / L
			end = end + count1*L
			count += count1
		}

	}

	return count
}

func main() {
	intervals := []Interval{}
	numRoad := 0
	lengthMu := 0
	a := 0
	b := 0
	fmt.Scan(&numRoad) //遇到换行符和空格都停下然后跳过读取
	fmt.Scan(&lengthMu)
	for i := 0; i < numRoad; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		intervals = append(intervals, Interval{start: a, end: b})
	}

	fmt.Println(coverMuddyRoad(numRoad, lengthMu, intervals))
}
