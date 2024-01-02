package main

import (
	"fmt"
	"strings"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

// 根据层次遍历和中序遍历构建二叉树
func buildTree(inorder []int, levelorder []int) *Node {
	if len(inorder) != 0 {
		for _, val := range levelorder {
			for i := range inorder {
				if inorder[i] == val {
					node := &Node{val: val}
					node.left = buildTree(inorder[:i], levelorder)
					node.right = buildTree(inorder[i+1:], levelorder)
					return node
				}
			}
		}
	}
	return nil
}

func preOrder(node *Node, result *[]int) {
	if node != nil {
		*result = append(*result, node.val)
		preOrder(node.left, result)
		preOrder(node.right, result)
	}
}

func main() {
	var num int
	fmt.Scan(&num)
	inorder := make([]int, num)
	levelorder := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&inorder[i])
	}
	for i := 0; i < num; i++ {
		fmt.Scan(&levelorder[i])
	}

	root := buildTree(inorder, levelorder)

	var result []int
	preOrder(root, &result)

	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]"))
}
