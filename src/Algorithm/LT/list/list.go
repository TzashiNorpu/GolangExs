package list

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	size int
	node *ListNode
}

type ListNode struct {
	Val  int
	next *ListNode
}

func GenLinkedListByInts(ints []int) *LinkedList {
	return &LinkedList{
		size: len(ints),
		node: genLinkedListByInts(ints, 0),
	}
}
func genLinkedListByInts(ints []int, level int) *ListNode {
	if level == len(ints)-1 {
		return &ListNode{Val: ints[level]}
	}
	curr := ListNode{Val: ints[level]}
	curr.next = genLinkedListByInts(ints, level+1)
	return &curr
}

func DelListNode(curr *ListNode, delNode *ListNode) {
	curr.next = delNode.next
}

func (l *LinkedList) AddFirst(n ListNode) *LinkedList {
	dummyHead := ListNode{next: l.node}
	n.next = dummyHead.next
	dummyHead.next = &n
	l.node = dummyHead.next
	l.size++
	return l
}

func (l *LinkedList) AddNth(index int, n ListNode) *LinkedList {
	if index < 0 || index > l.size {
		panic("Index out of range")
	}
	dummyHead := &ListNode{next: l.node}
	prev := dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	n.next = prev.next
	prev.next = &n
	l.size++
	return l
}

func (l *LinkedList) Traverse() {
	s := ""
	dummyHead := &ListNode{next: l.node}
	for dummyHead.next != nil {
		s = s + "->" + strconv.Itoa(dummyHead.next.Val)
		dummyHead.next = dummyHead.next.next
	}
	fmt.Println(s + "->nil")
}

func (l *LinkedList) DelNthNode(index int) {
	for {

	}
}

func (l *LinkedList) DelSpecificValueNode(node ListNode) *LinkedList {
	dummyHead := &ListNode{next: l.node}
	curr := dummyHead.next
	for curr != nil {
		if curr.Val == node.Val {
			dummyHead.next = curr.next
			curr = curr.next
			l.size--
		} else {
			curr = curr.next
		}
	}
	return &LinkedList{node: dummyHead.next}
}
