package lt61

import "golang/Algorithm/LT/list"

func rotateRight(head *list.ListNode, k int) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &list.ListNode{Next: head}

	size := 1
	curr := head
	for curr.Next != nil {
		size++
		curr = curr.Next
	}

	k = k % size
	if k == 0 {
		return head
	}

	curr.Next = dummyHead.Next
	prev := head
	for i := 0; i < size-k-1; i++ {
		prev = prev.Next
	}
	n := prev.Next
	prev.Next = nil
	return n
}
