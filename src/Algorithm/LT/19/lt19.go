package lt19

import "golang/Algorithm/LT/list"

func removeNthFromEnd(head *list.ListNode, n int) *list.ListNode {
	if head == nil {
		return nil
	}
	size := 0
	cntP := head
	dummyHead := &list.ListNode{Next: head}
	delP := dummyHead
	for ; cntP != nil; cntP = cntP.Next {
		size++
	}
	for i := 0; i < size-n; i++ {
		delP = delP.Next
	}
	delP.Next = delP.Next.Next
	return dummyHead.Next
}
