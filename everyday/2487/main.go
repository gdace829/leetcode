package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes1(head *ListNode) *ListNode {
	var nums []*ListNode
	for ; head != nil; head = head.Next {
		nums = append(nums, head)
	}
	var res *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		if res == nil {
			res = nums[i]
			nums[i].Next = nil
			continue
		}
		if nums[i].Val >= res.Val {
			nums[i].Next = res
			res = nums[i]
		}
	}

	return res
}

//	func reverse(head *ListNode) *ListNode {
//		var dummy *ListNode
//		for head != nil {
//			p := head
//			head = head.Next
//			p.Next = dummy
//			dummy = p
//		}
//		return dummy
//	}
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		p := head
		head = head.Next
		p.Next = pre
		pre = p
	}
	return pre
}
func removeNodes(head *ListNode) *ListNode {
	head = reverse(head)
	for p := head; p.Next != nil; {
		if p.Val > p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return reverse(head)
}
