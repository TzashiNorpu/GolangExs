package lt206

import "golang/Algorithm/LT/list"

func reverseList(head *list.ListNode) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reverse(head)
	return last
}

var last *list.ListNode

func reverse(head *list.ListNode) *list.ListNode {
	if head.Next == nil {
		last = head
		return head
	}
	if head.Next != nil {
		node := reverse(head.Next)
		head.Next = nil
		node.Next = head
	}
	return head
}

func reverseList2(head *list.ListNode) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	// 每一层的翻转
	head.Next.Next = head
	head.Next = nil
	// newHead 是从最后那一层带上来的   其他层没有修改
	return newHead
}
