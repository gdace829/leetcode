package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {

	sum := 0
	if root.Left != nil && root.Val > low {
		sum += rangeSumBST(root.Left, low, high)
	}
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Right != nil && root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}
	return sum
}
