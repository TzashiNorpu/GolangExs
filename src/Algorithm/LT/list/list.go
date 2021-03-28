package list

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	size int
	Node *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenLinkedListByInts(ints []int) *LinkedList {
	return &LinkedList{
		size: len(ints),
		Node: genLinkedListByInts(ints, 0),
	}
}
func genLinkedListByInts(ints []int, level int) *ListNode {
	if level == len(ints)-1 {
		return &ListNode{Val: ints[level]}
	}
	curr := ListNode{Val: ints[level]}
	curr.Next = genLinkedListByInts(ints, level+1)
	return &curr
}

func DelListNode(curr *ListNode, delNode *ListNode) {
	curr.Next = delNode.Next
}

func (l *LinkedList) AddFirst(n ListNode) *LinkedList {
	dummyHead := ListNode{Next: l.Node}
	n.Next = dummyHead.Next
	dummyHead.Next = &n
	l.Node = dummyHead.Next
	l.size++
	return l
}

func (l *LinkedList) AddNth(index int, n ListNode) *LinkedList {
	if index < 0 || index > l.size {
		panic("Index out of range")
	}
	dummyHead := &ListNode{Next: l.Node}
	prev := dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	n.Next = prev.Next
	prev.Next = &n
	l.size++
	return l
}

func (l *LinkedList) Traverse() {
	s := ""
	dummyHead := &ListNode{Next: l.Node}
	for dummyHead.Next != nil {
		s = s + strconv.Itoa(dummyHead.Next.Val) + "->"
		dummyHead.Next = dummyHead.Next.Next
	}
	fmt.Println(s + "nil")
}

func (l *LinkedList) DelNthNode(index int) *LinkedList {
	dummyHead := &ListNode{Next: l.Node}
	prev := dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	delNode := prev.Next
	prev.Next = delNode.Next
	delNode.Next = nil
	l.size--
	return &LinkedList{Node: dummyHead.Next, size: l.size}
}

func (l *LinkedList) DelSpecificValueNode(node ListNode) *LinkedList {
	head := l.Node
	size := l.size
	for head != nil && head.Val == node.Val {
		head = head.Next
		size--
	}
	if head == nil {
		return &LinkedList{Node: nil, size: 0}
	}
	prev := head
	for prev.Next != nil {
		if prev.Next.Val == node.Val {
			prev.Next = prev.Next.Next
			size--
		} else {
			prev = prev.Next
		}
	}

	return &LinkedList{Node: head, size: size}
}
