package lt203

import "golang/Algorithm/LT/list"

func removeElements(head *list.ListNode, val int) *list.ListNode {
	if head == nil {
		return head
	}
	prev := &list.ListNode{Next: head}
	for prev.Next != nil && prev.Next.Val == val {
		prev.Next = prev.Next.Next
		head = head.Next
	}

	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}

func removeElements2(head *list.ListNode, val int) *list.ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		// 当前节点值和val相同返回当前节点的下一个节点
		// 递归返回上一层时将上一层的节点和当前节点的下一个节点关联[跳过当前节点]
		return head.Next
	} else {
		return head
	}
}
