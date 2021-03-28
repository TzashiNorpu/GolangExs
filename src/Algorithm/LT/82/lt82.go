package lt82

import "golang/Algorithm/LT/list"

func deleteDuplicates(head *list.ListNode) *list.ListNode {
	if head == nil {
		return head
	}
	curr := head
	m := make(map[int]int)
	for curr != nil {
		m[curr.Val] = m[curr.Val] + 1
		curr = curr.Next
	}

	for head != nil && m[head.Val] > 1 {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	prev := head
	for prev.Next != nil {
		if prev.Next != nil && m[prev.Next.Val] > 1 {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}

func deleteDuplicates2(head *list.ListNode) *list.ListNode {
	if head == nil {
		return head
	}
	dummyHead := &list.ListNode{Next: head}
	curr := dummyHead
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			x := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == x {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return dummyHead.Next
}
