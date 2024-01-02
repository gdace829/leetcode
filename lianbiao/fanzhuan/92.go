package fanzhuan

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	shaobin := &ListNode{}
	shaobin.Next = head
	start := shaobin
	times := right - left + 1
	for left-1 != 0 {
		start = start.Next
		left--
	}
	pre := start
	cur := start.Next
	for times != 0 {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		times--
	}
	start.Next.Next = cur
	start.Next = pre

	return shaobin.Next
}
