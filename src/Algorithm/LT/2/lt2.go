package lt2

import (
	"golang/Algorithm/LT/list"
)

var sum int = 0

func addTwoNumbers(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	l1Size, l2Size := 1, 1
	head := l1
	for head.Next != nil {
		l1Size++
		head = head.Next
	}
	head = l2
	for head.Next != nil {
		l2Size++
		head = head.Next
	}
	//1->9->2->3
	//      9->9
	if l1Size < l2Size {
		l1, l2 = l2, l1
		l1Size, l2Size = l2Size, l1Size
	}
	s := l2
	for i := 0; i < l2Size-1; i++ {
		s = s.Next
	}
	for i := 0; i < l1Size-l2Size; i++ {
		s.Next = &list.ListNode{Next: nil}
		s = s.Next
	}
	padding := l2
	res := &list.ListNode{Next: nil}
	track := res
	step := 0
	for i := 0; i < l1Size; i++ {
		val := (l1.Val + padding.Val + step) % 10
		track.Next = &list.ListNode{Val: val, Next: nil}
		if l1 != nil {
			step = (l1.Val + padding.Val + step) / 10
		}
		l1 = l1.Next
		padding = padding.Next
		track = track.Next
	}
	if step >= 1 {
		track.Next = &list.ListNode{Val: step}
	}
	return res.Next
}

func addTwoNumbers2(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	p1, p2 := l1, l2
	step := 0
	p := &list.ListNode{Next: nil}
	res := p
	for p1 != nil || p2 != nil {
		v1, v2 := 0, 0
		if p1 != nil {
			v1 = p1.Val
		}
		if p2 != nil {
			v2 = p2.Val
		}
		val := (v1 + v2 + step) % 10
		step = (v1 + v2 + step) / 10
		p.Next = &list.ListNode{Val: val, Next: nil}
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}

		p = p.Next
	}
	if step == 1 {
		p.Next = &list.ListNode{Val: 1}
	}
	return res.Next
}
