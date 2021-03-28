package lt83

import "golang/Algorithm/LT/list"

func deleteDuplicates(head *list.ListNode) *list.ListNode {
	if head == nil {
		return nil
	}
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
