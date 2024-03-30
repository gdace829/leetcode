package main

import "math/bits"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	return FindElements{root: root}
}

func (this *FindElements) Find(target int) bool {
	target++
	cur := this.root
	for i := bits.Len(uint(target)) - 2; i >= 0; i-- {
		bit := target >> i & 1
		if bit == 0 {
			cur = cur.Left
		}
		if bit == 1 {
			cur = cur.Right
		}
		if cur == nil {
			return false
		}
	}
	return true
}
