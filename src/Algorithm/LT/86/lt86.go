package lt86

import "golang/Algorithm/LT/list"

func partition(head *list.ListNode, x int) *list.ListNode {
	dummyHead1 := &list.ListNode{Next: nil}
	l1 := dummyHead1
	dummyHead2 := &list.ListNode{Next: nil}
	l2 := dummyHead2
	for head != nil {
		if head.Val < x {
			dummyHead1.Next = head
			dummyHead1 = dummyHead1.Next
		} else {
			dummyHead2.Next = head
			dummyHead2 = dummyHead2.Next
		}
		head = head.Next
	}
	if dummyHead1.Next != nil {
		dummyHead1.Next = nil
	}
	if dummyHead2.Next != nil {
		dummyHead2.Next = nil
	}
	dummyHead1.Next = l2.Next
	return l1.Next
}
