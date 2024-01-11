package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 求两个数的最大公约数a,b,a%b的公约数相同
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	node := head
	for node.Next != nil {
		node.Next = &ListNode{gcd(node.Val, node.Next.Val), node.Next}
		node = node.Next.Next
	}
	return head
}
func main() {

}
