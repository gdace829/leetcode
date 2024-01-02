package fanzhuan
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	shaobin := &ListNode{}
	res := shaobin
	shaobin.Next = head
	num := 0
	//统计链表数目
	for head != nil {
		head = head.Next
		num++
	}
	pre := shaobin
	cur := shaobin.Next
	for num/k > 0 {
		k1 := k
		//转换k个
		for k1 != 0 {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
			k1--
		}
		shaobin1 := shaobin.Next
		shaobin.Next.Next = cur
		shaobin.Next = pre
		shaobin = shaobin1
		num = num - k
	}
	return res.Next
}
