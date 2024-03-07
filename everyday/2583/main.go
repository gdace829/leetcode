package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	sumLevel := []int{}
	for len(q) > 0 {
		size := len(q)
		sum := 0
		var node *TreeNode
		for i := 0; i < size; i++ {
			node = q[0]
			sum += node.Val
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		sumLevel = append(sumLevel, sum)
	}
	sort.Slice(sumLevel, func(i, j int) bool { return sumLevel[i] > sumLevel[j] })
	if len(sumLevel) < k {
		return -1
	}
	return int64(sumLevel[k-1])
}
