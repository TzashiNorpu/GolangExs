package test1

import (
	"fmt"
	"golang/Algorithm/list"
)

func deleteDuplicates(head *list.ListNode) *list.ListNode {
	dummyHead1 := &list.ListNode{Next: head}
	m := make(map[int]int)
	for dummyHead1.Next != nil {
		m[dummyHead1.Next.Val] = m[dummyHead1.Next.Val] + 1
		dummyHead1 = dummyHead1.Next
	}
	for k, v := range m {
		fmt.Printf("key=%d,value=%d\n", k, v)
	}
	dummyHead2 := &list.ListNode{Next: head}
	for dummyHead2.Next != nil {
		if m[dummyHead2.Next.Val] > 1 {
			dummyHead2.Next = dummyHead2.Next.Next
		}
		dummyHead2 = dummyHead2.Next
	}

	return dummyHead2
}
