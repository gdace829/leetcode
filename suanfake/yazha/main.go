package main

import (
	"fmt"
	"sort"
)

type Job struct {
	start int
	end   int
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxWorkTime(jobs []Job) int {
	var m int
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].end == jobs[j].end {
			return jobs[i].start < jobs[j].start
		}
		return jobs[i].end < jobs[j].end
	})
	dp := make([]int, len(jobs))
	for i := 0; i < len(jobs); i++ {
		dp[i] = jobs[i].end - jobs[i].start
	}
	for i := 1; i < len(jobs); i++ {
		for j := 0; j < i; j++ {
			if jobs[j].end <= jobs[i].start {
				dp[i] = max(dp[j]+jobs[i].end-jobs[i].start, dp[i])
			}
			if dp[i] > m {
				m = dp[i]
			}
		}
	}

	return m
}

func main() {

	jobs := []Job{}
	var n, a, b int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		jobs = append(jobs, Job{})

	}
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		jobs[i].start = a

	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		jobs[i].end = b

	}
	fmt.Println(maxWorkTime(jobs))
}
