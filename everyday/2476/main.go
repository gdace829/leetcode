package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	list := []int{}
	var zhongxubianli func(*TreeNode)
	ans := [][]int{}
	zhongxubianli = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		zhongxubianli(tn.Left)
		list = append(list, tn.Val)
		zhongxubianli(tn.Right)
	}
	for i := 0; i < len(queries); i++ {

	}
	val := queries[0]
	left := 0
	right := len(list) - 1
	for left < right {
		mid := (left + right) / 2
		if list[mid] == val {
			ans = append(ans, []int{val, val})
			break
		}
		if true {

		}
	}
	return ans
}
