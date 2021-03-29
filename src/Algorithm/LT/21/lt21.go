package lt21

import "golang/Algorithm/LT/list"

func mergeTwoLists(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	prev := &list.ListNode{Next: nil}
	dummyHead := prev
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return dummyHead.Next
}

//list1[0]+merge(list1[1:],list2) list1[0]<list2[0]
//list2[0]+merge(list1,list2[1:]) otherwise
// 3->3->4
// 1->2->5

func mergeTwoLists2(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}
