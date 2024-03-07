package main

import "slices"

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

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	tree := &TreeNode{Val: preorder[0]}
	index := slices.Index(postorder, preorder[1]) + 1
	tree.Left = constructFromPrePost(preorder[1:1+index], postorder[:index])
	tree.Right = constructFromPrePost(preorder[1+index:], postorder[index:len(postorder)-1])
	return tree
}
