package main

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	idMap := make(map[int]int)
	for k, v := range inorder {
		idMap[v] = k
	}
	var build func(int, int) *TreeNode
	build = func(i1, i2 int) *TreeNode {
		if i1 > i2 {
			return nil
		}
		val := postorder[len(postorder)-1]
		tree := &TreeNode{Val: val}
		postorder = postorder[:len(postorder)-1]
		tree.Right = build(idMap[val]+1, i2)
		tree.Left = build(i1, idMap[val]-1)
		return tree
	}
	return build(0, len(inorder)-1)
}
