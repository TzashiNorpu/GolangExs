package lt24

import "golang/Algorithm/LT/list"

func swapPairs(head *list.ListNode) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &list.ListNode{Next: head}
	prev := dummyHead
	var first *list.ListNode
	var second *list.ListNode
	for prev.Next != nil && prev.Next.Next != nil {
		first = prev.Next
		prev.Next = prev.Next.Next
		second = prev.Next.Next
		prev.Next.Next = first
		first.Next = second
		prev = first
	}
	return dummyHead.Next
}
